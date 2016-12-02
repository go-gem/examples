package main

import (
	"fmt"

	"github.com/go-gem/gem"
	"github.com/go-gem/log"
)

type user struct {
	Name string `json:"name"`
}

var (
	u = user{Name: "foo"}
)

func html(ctx *gem.Context) {
	ctx.HTML(200, "hello world.")
}

func json(ctx *gem.Context) {
	ctx.JSON(200, u)
}

func jsonp(ctx *gem.Context) {
	ctx.JSONP(200, u, []byte("success"))
}

func xml(ctx *gem.Context) {
	ctx.XML(200, u)
}

func profile(ctx *gem.Context) {
	ctx.HTML(200, fmt.Sprintf("hello %s.", ctx.Param("name")))
}

func list(ctx *gem.Context) {
	ctx.Logger().Printf("page number: %d\n", ctx.ParamInt("page"))
	ctx.HTML(200, fmt.Sprintf("page number %s", ctx.Param("page")))
}

func main() {
	router := gem.NewRouter()

	router.GET("/", html)
	router.GET("/json", json)
	router.GET("/jsonp", jsonp)
	router.GET("/xml", xml)
	router.GET("/profile/:name", profile)
	router.GET("/list/:page", list)

	srv := gem.New(":8080", router.Handler())

	log.Fatal(srv.ListenAndServe())
}
