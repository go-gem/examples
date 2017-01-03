package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/go-gem/gem"
)

// index handler function.
func index(ctx *gem.Context) {
	ctx.HTML(200, "hello world")
}

func main() {
	// Create server.
	srv := gem.New(":8080")

	srv.SetLogger(logrus.New())

	// Create router.
	router := gem.NewRouter()

	// Use debug middleware
	router.Use(&Debug{})

	// Register index handler
	router.GET("/", index)

	router.GET("/logger", func(ctx *gem.Context) {
		ctx.Logger().Debug("debug")
		ctx.Logger().Info("info")
		ctx.Logger().Error("error")
	})

	// Static files.
	router.ServeFiles("/tmp/*filepath", http.Dir(os.TempDir()))

	// REST APIs.
	router.GET("/users", userList)
	router.POST("/users", userAdd)
	router.GET("/users/:name", userProfile)
	router.PUT("/users/:name", userUpdate)
	router.DELETE("/users/:name", userProfile)

	// Start server.
	log.Println(srv.ListenAndServe(router.Handler()))
}

func userList(ctx *gem.Context) {

}

func userAdd(ctx *gem.Context) {
	ctx.Request.ParseForm()
	ctx.FormValue("name")
	ctx.Request.ParseForm()
	ctx.FormValue("name")

}

func userProfile(ctx *gem.Context) {

}

func userUpdate(ctx *gem.Context) {

}

type Debug struct{}

func (d *Debug) Wrap(next gem.Handler) gem.Handler {
	return gem.HandlerFunc(func(ctx *gem.Context) {
		// print request info.
		log.Println(ctx.Request.URL, ctx.Request.Method)

		// call the next handler.
		next.Handle(ctx)
	})
}
