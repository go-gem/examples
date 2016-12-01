package main

import (
	"fmt"

	"github.com/go-gem/gem"
	"github.com/go-gem/log"
)

func main() {
	router := gem.NewRouter()

	router.GET("/user/:name", func(ctx *gem.Context) {
		ctx.HTML(200, fmt.Sprintf("Hello %s", ctx.Param("name")))
	})

	log.Fatal(gem.ListenAndServe(":8080", router.Handler()))
}
