# Logger

Gem defines a [Logger](https://godoc.org/github.com/go-gem/gem#Logger) interface, so it is easy to custom logger.

AFAIK, the following logging packages are compatible with Gem:

- [gem log](https://github.com/go-gem/log) - a simple and leveled logging package, maintained by Gem.
- [logrus](https://github.com/sirupsen/logrus) - structured, pluggable logging package.

Run the logger's example:

```
go run $GOPATH/src/github.com/go-gem/examples/logger/server.go
```

And then navigate to [localhost:8080](http://localhost:8080):

```
Hello world.
```

Console output:

```
INFO[0002] 1                                            
INFO[0002] 2                                            
WARN[0002] 3                                            
ERRO[0002] 4
```