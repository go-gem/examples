package main

import (
	"os"

	"github.com/go-gem/gem"
	"github.com/go-gem/log"
)

func main() {
	router := gem.NewRouter()

	// GET
	router.GET("/", func(ctx *gem.Context) {
		ctx.HTML(200, "hello world.")
	})

	// Serve static files.
	router.ServeFiles("/tmp/*filepath", os.TempDir())

	srv := gem.New(":8080", router.Handler())

	log.Println(srv.ListenAndServe())
}
