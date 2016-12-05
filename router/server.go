package main

import (
	"os"
	"sync"

	"github.com/go-gem/gem"
	"github.com/go-gem/log"
	"github.com/valyala/fasthttp"
)

type debugMiddleware struct{}

func (m *debugMiddleware) Handle(next gem.Handler) gem.Handler {
	return gem.HandlerFunc(func(ctx *gem.Context) {
		// print the request's method and path.
		ctx.Logger().Printf("%s: %s\n", ctx.Method(), ctx.Path())

		next.Handle(ctx)
	})
}

type user struct {
	Name    string `json:"name"`
	Company string `json:"company"`
}

var (
	mutex = sync.RWMutex{}
	users = map[string]*user{
		"foo": &user{Name: "foo", Company: "comfoo"},
		"bar": &user{Name: "bar", Company: "combar"},
	}
)

func userList(ctx *gem.Context) {
	ctx.JSON(200, users)
}

func userInfo(ctx *gem.Context) {
	// get the user's name.
	name := ctx.Param("name")

	mutex.RLock()
	defer mutex.RUnlock()
	if user, ok := users[name]; ok {
		ctx.JSON(200, user)
		return
	}

	ctx.JSON(fasthttp.StatusNotFound, nil)
}

func userAdd(ctx *gem.Context) {
	name := string(ctx.FormValue("name"))
	company := string(ctx.FormValue("company"))

	mutex.Lock()
	defer mutex.Unlock()
	user := &user{Name: name, Company: company}
	users[name] = user
	ctx.JSON(200, user)
}

func userDelete(ctx *gem.Context) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(users, ctx.Param("name"))
	ctx.JSON(fasthttp.StatusOK, nil)
}

func userUpdate(ctx *gem.Context) {
	name := ctx.Param("name")

	mutex.RLock()
	defer mutex.RUnlock()
	if user, ok := users[name]; ok {
		user.Company = string(ctx.FormValue("company"))
		ctx.JSON(fasthttp.StatusOK, user)
		return
	}

	ctx.JSON(fasthttp.StatusNotFound, nil)
}

func main() {
	router := gem.NewRouter()

	// Use middleware.
	router.Use(&debugMiddleware{})

	// GET
	router.GET("/", func(ctx *gem.Context) {
		ctx.HTML(200, "hello world.")
	})

	// Serve static files.
	router.ServeFiles("/tmp/*filepath", os.TempDir())

	// REST APIs
	router.GET("/users", userList)
	router.GET("/users/:name", userInfo)
	router.POST("/users", userAdd)
	router.PUT("/users/:name", userUpdate)
	router.DELETE("/users/:name", userDelete)

	// Create a server instance.
	srv := gem.New(":8080", router.Handler())

	log.Println(srv.ListenAndServe())
}
