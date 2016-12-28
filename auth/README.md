# HTTP Basic and Digest Authentication Middleware for [Gem](https://github.com/go-gem/gem) Web Framework

**Start server**

```
$ go run $GOPATH/src/github.com/go-gem/examples/auth/server.go
```

And then navigate to [Basic Authentication](localhost:8080/basic) and [Digest Authentication](localhost:8080/degest).
Username: `foo`, password: `bar`.

This package depends on [go-http-auth](https://github.com/abbot/go-http-auth), more usages may be obtained on
https://github.com/abbot/go-http-auth.