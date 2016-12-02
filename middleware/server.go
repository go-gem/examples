package main

import "github.com/go-gem/gem"

func main() {

}

// MyMiddleware my first middleware
type MyMiddleware struct {
}

// Handle implements the Middleware's Handle function.
func (m *MyMiddleware) Handle(next gem.Handler) gem.Handler {
	return gem.HandlerFunc(func(ctx *gem.Context) {

	})
}
