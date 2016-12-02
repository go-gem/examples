# Context

[Wiki](https://github.com/go-gem/gem/wiki/Context)

Start the server:

```
go run $GOPATH/src/github.com/go-gem/examples/context/server.go
```

And then navigate to:

1. [http://localhost:8080](http://localhost:8080)

```
hello world.
```

2. [http://localhost:8080/json](http://localhost:8080/json)

```
{"name":"foo"}
```

3. [http://localhost:8080/jsonp](http://localhost:8080/jsonp)

```
success({"name":"foo"})
```

4. [http://localhost:8080](http://localhost:8080)

```
<?xml version="1.0" encoding="UTF-8"?>
<user><Name>foo</Name></user>
```

5. [http://localhost:8080/profile/bar](http://localhost:8080/profile/bar)

```
hello bar.
```

5. [http://localhost:8080/list/2](http://localhost:8080/list/2)
```
page number 2
```
