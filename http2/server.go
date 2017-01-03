package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/go-gem/gem"
)

var (
	curDir   = path.Join(os.Getenv("GOPATH"), "src", "github.com", "go-gem", "examples", "http2")
	certFile = path.Join(curDir, "cert.pem")
	keyFile  = path.Join(curDir, "key.pem")
)

func main() {
	srv := gem.New(":8080")

	router := gem.NewRouter()
	router.GET("/", func(ctx *gem.Context) {
		if err := ctx.Push("/images/logo.png", nil); err != nil {
			ctx.Logger().Info(err)
		}

		ctx.HTML(200, `<html><head></head><body><img src="/images/logo.png"/></body></html>`)
	})
	router.ServeFiles("/images/*filepath", http.Dir(path.Join(curDir, "images")))

	log.Println(srv.ListenAndServeTLS(certFile, keyFile, router.Handler()))
}
