package main

import (
	"log"
	"os"

	"github.com/go-gem/gem"
)

func handler(ctx *gem.Context) {
	ctx.HTML(200, "Hello world.")
}

func newHandler(ctx *gem.Context) {
	ctx.HTML(200, "Congratulation! Server has been upgraded.")
}

func main() {
	log.Printf("Server started at process %d", os.Getpid())

	srv := gem.New(":8080", handler)

	log.Print(srv.ListenAndServe())
}
