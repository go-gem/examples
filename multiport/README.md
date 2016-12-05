# Listen multiple ports at single process

This example shows that how to listen multiple ports at single process.

**Complie and start the server**

```
cd $GOPATH/src/github.com/go-gem/examples/multiport
go install
$GOPATH/bin/multiport
```

**Output**
```
2016/12/05 14:49:32 process id: 5319
```

And then let's navigate to [localhost:8080](http://localhost:8080) and [localhost:8081](http://localhost:8081)

| Url                                   | Output     |
|:--------------------------------------|: ----------|
|[localhost:8080](http://localhost:8080)| Hello foo. |
|[localhost:8081](http://localhost:8081)| Hello bar. |