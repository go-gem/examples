# Sessions

[Wiki](https://github.com/go-gem/gem/wiki/Sessions)

A simple example of sign in and sign out.

Start the server:

```
go run $GOPATH/src/github.com/go-gem/examples/session/server.go
```

And then navigate to [http://localhost:8080](http://localhost:8080).

**Important note**: it is **necessary** that use the sessions middleware to save and clear session,
otherwise, any changes of session would not be save, not only that, all of sessions are confused,
because the `fasthttp` using sync.Pool to reuse `fasthttp.RequestCtx` instance.

```
router.Use(middleware.NewSessions())
```