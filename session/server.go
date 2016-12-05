package main

import (
	"html/template"

	"github.com/go-gem/gem"
	"github.com/go-gem/gem/middleware"
	"github.com/go-gem/log"
	"github.com/go-gem/sessions"
	"github.com/valyala/fasthttp"
)

var (
	tplIndex = template.Must(template.New("index").Parse(`<html><head><title>Welcome</title></head><body>
	Hello {{.}}, <a href="/sign-out">click here to signout</a>
	</body></html>`))

	tplSignIn = template.Must(template.New("signIn").Parse(`<html>
	<head><title>Sign in</title></head><body>
	<form action="/sign-in" method="post">
	<input name="name" placeholder="username">
	<button type="submit">Sign in</button>
	</form>
	</body></html>`))

	store = sessions.NewCookieStore([]byte("something-very-secret"))
)

func getSession(ctx *gem.Context) *sessions.Session {
	// session always is non-nil, even if an non-nil error was returned,
	// you can ignore the error, but print the error would be useful for
	// debugging.
	session, err := ctx.SessionsStore().Get(ctx.RequestCtx, "_user")
	defer session.Save(ctx.RequestCtx)
	if err != nil {
		ctx.Logger().Debugln(err)
	}

	return session
}

func index(ctx *gem.Context) {
	ctx.SetContentType(gem.HeaderContentTypeHTML)

	session := getSession(ctx)

	var name string
	var ok bool
	if name, ok = session.Values["name"].(string); !ok || name == "" {
		// guest
		tplSignIn.Execute(ctx, nil)
		return
	}

	tplIndex.Execute(ctx, name)
}

func signIn(ctx *gem.Context) {
	name := string(ctx.PostArgs().Peek("name"))
	if name == "" {
		ctx.HTML(200, "the name can not be blank.")
		return
	}

	// get session and remember the username.
	session := getSession(ctx)
	session.Values["name"] = name

	ctx.Redirect("/", fasthttp.StatusOK)
}

func signOut(ctx *gem.Context) {
	// clear the username from session.
	session := getSession(ctx)
	delete(session.Values, "name")

	ctx.Redirect("/", fasthttp.StatusOK)
}

func main() {
	router := gem.NewRouter()

	// It is necessary that use the sessions middleware to
	// save and clear session.
	router.Use(middleware.NewSessions())

	router.GET("/", index)
	router.POST("/sign-in", signIn)
	router.GET("/sign-out", signOut)

	srv := gem.New(":8080", router.Handler())

	srv.SetSessionsStore(store)

	log.Fatal(srv.ListenAndServe())
}
