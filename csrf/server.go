package main

import (
	"log"
	"html/template"
	"fmt"

	"github.com/go-gem/gem"
	"github.com/go-gem/gem/middleware"
)

var (
	users = map[string]string{
		"foo":"foopsw",
		"bar":"barpsw",
	}
)

// csrf configuration.
var (
	csrfMiddleware = middleware.NewCSRF()
)

var (
	tpl = template.New("login")
)

func loginGet(ctx *gem.Context) {
	ctx.Response.Header.SetContentTypeBytes(gem.ContentTypeHTML)
	err := tpl.Execute(ctx, map[string]interface{}{
		"CSRF":ctx.UserValue("_csrf"),
	})

	if err != nil {
		fmt.Printf(err.Error())
	}
}

func loginPost(ctx *gem.Context) {
	name := string(ctx.PostArgs().Peek("name"))
	psw := string(ctx.PostArgs().Peek("psw"))

	if truePsw, ok := users[name]; ok && psw == truePsw {
		ctx.HTML(200, "Login successfully.")
		return
	}

	ctx.HTML(200, "Incorrect username or password.")
}

func main() {
	var err error
	if tpl, err = tpl.Parse(`
	<form method="post" action="/login">
		<input type="hidden" name="_csrf" value="{{.CSRF}}">

		<label>username</label>
		<input name="name">

		<label>password</label>
		<input type="password" name="psw">

		<button type="submit">Login</button>

	</form>
	`); err != nil {
		panic(err)
	}

	router := gem.NewRouter()

	router.Use(csrfMiddleware)

	router.GET("/login", loginGet)

	router.POST("/login", loginPost)

	log.Fatal(gem.ListenAndServe(":1234", router.Handler))
}


