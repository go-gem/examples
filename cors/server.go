package main

import (
	"github.com/go-gem/gem"
	"github.com/go-gem/middleware-cors"
	"github.com/rs/cors"
)

var corsMiddleware = corsmidware.New(cors.Options{})

func main() {
	router := gem.NewRouter()
	router.Use(corsMiddleware)
	router.GET("/", func(ctx *gem.Context) {
		ctx.JSON(200, user{Name: "foo"})
	})

	gem.ListenAndServe(":8080", router.Handler())
}

type user struct {
	Name string `json:"name"`
}
