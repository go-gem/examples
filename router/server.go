package main

import (
	"os"

	"github.com/go-gem/gem"
	"github.com/go-gem/log"
)

type debugMiddleware struct{}

func (m *debugMiddleware) Handle(next gem.Handler) gem.Handler {
	return gem.HandlerFunc(func(ctx *gem.Context) {
		// print the request's method and path.
		ctx.Logger().Printf("%s: %s\n", ctx.Method(), ctx.Path())

		next.Handle(ctx)
	})
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

	srv := gem.New(":8080", router.Handler())

	log.Println(srv.ListenAndServe())
}
