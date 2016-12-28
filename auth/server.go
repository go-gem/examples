package main

import (
	"github.com/abbot/go-http-auth"
	"github.com/go-gem/gem"
	"github.com/go-gem/middleware-auth"
)

var (
	salt     = []byte("salt")
	magic    = []byte("$1$")
	username = "foo"
	password = []byte("bar")
)

// basic auth middleware
var (
	basicPasswd = string(auth.MD5Crypt(password, salt, magic))

	basicAuthenticator = auth.NewBasicAuthenticator("", func(user, realm string) string {
		if user == username {
			return basicPasswd
		}
		return ""
	})

	basicAuth = authmidware.New(basicAuthenticator)
)

func basicHandle(ctx *gem.Context) {
	ctx.HTML(200, "hello "+basicAuth.Username(ctx))
}

// digest auth middleware
var (
	digestAuthenticator = auth.NewDigestAuthenticator("", func(user, realm string) string {
		if user == "foo" {
			// MD5(username:realm:password)
			return "e0a109b991367f513dfa73bbae05fb07"
		}
		return ""
	})

	digestAuth = authmidware.New(digestAuthenticator)
)

func digestHandle(ctx *gem.Context) {
	ctx.HTML(200, "hello "+digestAuth.Username(ctx))
}

func main() {
	router := gem.NewRouter()

	// basic auth
	router.GET("/basic", basicHandle, &gem.HandlerOption{
		Middlewares: []gem.Middleware{basicAuth},
	})

	// digest auth
	router.GET("/digest", digestHandle, &gem.HandlerOption{
		Middlewares: []gem.Middleware{digestAuth},
	})

	gem.ListenAndServe(":8080", router.Handler())
}
