package main

import (
	"os"
	"sync"

	"github.com/go-gem/gem"
	"github.com/go-gem/log"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		s := gem.New(":8080", func(ctx *gem.Context) {
			ctx.SetBodyString("Hello foo.")
		})
		log.Println(s.ListenAndServe().Error())
		wg.Done()
	}()

	go func() {
		s := gem.New(":8081", func(ctx *gem.Context) {
			ctx.SetBodyString("Hello bar.")
		})
		log.Println(s.ListenAndServe().Error())
		wg.Done()
	}()

	pid := os.Getpid()

	log.Println("process id:", pid)

	wg.Wait()
}
