# Router

[Wiki](https://github.com/go-gem/gem/wiki/Router)

Start the server:

```
go run $GOPATH/src/github.com/go-gem/examples/router/server.go
```

And then navigate to:

1. [http://localhost:8080](http://localhost:8080)

```
hello world.
```

2. [http://localhost:8080/tmp](http://localhost:8080/tmp)

```
...
files
...
```

Check out your console, some debug messages come in slight:

```
GET: /
GET: /tmp
```

**REST APIs**

|Url                                                                    | Method | Output                                                                          |
|:-----------------------------------------------------------------     |:------ |:--------------------------------------------------------------------------------|
|[/users](http://localhost:8080/users)                                  | GET    |{"bar":{"name":"bar","company":"combar"},"foo":{"name":"foo","company":"comfoo"}}|
|[/users/foo](http://localhost:8080/users/foo)                          | GET    |{"name":"foo","company":"comfoo"}                                                |
|[/users/foo](http://localhost:8080/users/foo)                          | DELETE |null                                                                             |
|[/users?name=a&company=b](http://localhost:8080/users?name=a&company=b)| POST   |{"name":"a","company":"b"}                                                       |
|[/users/a?company=b](http://localhost:8080/users/a?company=b)          | PUT    |{"name":"a","company":"new"}                                                     |
|[/users](http://localhost:8080/users)                                  | GET    |{"a":{"name":"a","company":"new"},"bar":{"name":"bar","company":"combar"}}       |

