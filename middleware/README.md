# Middleware

[Wiki](https://github.com/go-gem/gem/wiki/Middleware)

**Middleware** always play important roles in web server, such as **debugger**, **blocker**(filter) or **preprocessor**.

## How to write a middleware?

Let's take a look at [Middleware](https://godoc.org/github.com/go-gem/gem#Middleware) interface.

```
type Middleware interface {
    Handle(next Handler) Handler
}
```

It is very simple, our middleware just need implements the `Handle(next Handler) Handler` function.

Let's write our middleware step by step:

Firstly, defines an empty struct named `MyMiddleware`:

```
type MyMiddleware struct {}
```

Secondly, implements the Middleware's Handle function:

```
func (m *MyMiddleware) Handle(next Handler) Handler {
    return gem.HandlerFunc(func(ctx *gem.Context){
        
        // invoke the next handler
        next.Handle(ctx)
    })
} 
```


### Built-in middlewares
