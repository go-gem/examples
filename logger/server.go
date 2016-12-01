package main

import (
	"log"

	"github.com/Sirupsen/logrus"
	"github.com/go-gem/gem"
)

func main() {
	logger := logrus.New()

	srv := gem.New(":8080", func(ctx *gem.Context) {
		ctx.Logger().Println("1")
		ctx.Logger().Infoln("2")
		ctx.Logger().Warningln("3")
		ctx.Logger().Errorln("4")
		ctx.HTML(200, "Hello world.")
	})

	srv.SetLogger(logger)

	log.Fatal(srv.ListenAndServe())
}
