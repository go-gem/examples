# Graceful shutdown and restart

Graceful shutdown and restart is a very useful feature, it allow you to upgrade server or configuration,  and at the same time,
the old server wont be shutdown until all of in-process connections have been finished or reached timeout error.

## Install

```
go get github.com/go-gem/examples/graceful
```

## Restart

Let me show you how to restart your server gracefully.

### Step one

**Compile**

```
cd $GOPATH/src/github.com/go-gem/examples/graceful
go install
```

**Start server**

```
$GOPATH/bin/graceful
2016/12/01 19:48:53 Server started at process 29472
```

Navigate to [localhost:8080](http://localhost:8080)

```
Hello world.
```

### Step two

1. Firstly, change the server's `handler` to `newHandler`, and then recompile the server.

2. Secondly, send a `Restart` signal to the server process `29472`.

```
kill -HUP 2472
```

and you will received the following messages:

```
2016/12/01 20:26:59 [29472] received signal "hangup".
2016/12/01 20:26:59 [29472] Restarting...
2016/12/01 20:26:59 [29472] Fork-exec to 30979.
2016/12/01 20:26:59 Server started at process 30979
2016/12/01 20:26:59 [29472] received signal "terminated".
2016/12/01 20:26:59 [29472] shutting down the server(:8080)...
2016/12/01 20:26:59 [29472] server(:8080) shutdown successfully.
2016/12/01 20:26:59 [29472] all of old servers have been shutdown successfully.
```

3. Finally, let's navigate to [localhost:8080](http://localhost:8080):

```
Congratulation! Server has been upgraded.
```


## Shutdown

It is easy to shut down your server gracefully by running the following command:

```
kill -TERM 30979
```

```
2016/12/01 20:33:55 [30979] received signal "terminated".
2016/12/01 20:33:55 [30979] shutting down the server(:8080)...
2016/12/01 20:34:10 [30979] server(:8080) has been shutdown, but some exsiting connctions reach error: timeout.
2016/12/01 20:34:10 [30979] all of old servers have been shutdown successfully.
```

## Signals

| Signal          | Action   |
|:------          |:---------|
| syscall.SIGTERM | Shutdown |
| syscall.SIGHUP  | Restart  |
| the others      | Ignore   |

See also [SetSignal](https://godoc.org/github.com/go-gem/gem#SetSignal) to custom signal's action.

Note that: `syscall.SIGTERM` is not allow to custom!

## WaitTimeout

By default, server.waitTimeout is 15 seconds, you can change it by the API [Server.LoadConfig](https://godoc.org/github.com/go-gem/gem#Server.LoadConfig).

See also [ServerConfig](https://godoc.org/github.com/go-gem/gem#ServerConfig)