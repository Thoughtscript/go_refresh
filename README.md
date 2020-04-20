# golang_exm

Go Lang examples - some very simple examples.

*(More advanced examples combining these and other features are implemented elsewhere.)*

# How Do I Use This?

Navigate manually to one of the main root project directories (e.g. - `asyncGo`, `templateServer` and `httpServer`) using Bash:

```bash
$   go run httpServer.go
```

```bash
$   go run restApi.go
```

```bash
$   go run templateServer.go
```

# What's Here?

**httpServer** - this provides several simple http server functions using the native http server provided by GoLang.

Main endpoints:

```
    localhost:8888/print <- simple String to HTML writeout
    localhost:8888/write <- simple Model String to HTML writeout
    localhost:8888/public/index.html or localhost:8888/public/
```

**templateServer** - this provides several simple http server functions using the native http server provided by GoLang.

```
    localhost:7777/home <- simple String to HTML
```

# Gotcha's

To use the `syncRestApi` you must set the `GO_PATH` to `syncRestApi` in order for the default `vendor` location to be automatically injected into the app.

The other projects use their respective sub-directories (e.g. - `asyncGo` and `httpServer`).

Running the two server examples doesn't seem to work using the JetBrains Go IDE. This is probably user error (my fault) and due to the configuration settings or project layout I've selected.

As such, it's imperative to use the correct Bash commands per **How Do I Use This Above** the above.

## Fantastic Resources

1. REST: 
    * http://www.blog.labouardy.com/build-restful-api-in-go-and-mongodb/
    * Gin: https://gin-gonic.github.io/gin/
1. Go Routines: 
    * https://medium.com/@nikolay.bystritskiy/how-i-tried-to-do-things-asynchronously-in-golang-40e0c1a06a66
    * https://gobyexample.com/goroutines
1. Static: 
    * http://www.alexedwards.net/blog/serving-static-sites-with-go
    * https://rockfloat.com/post/learning-golang-templates.html
    * https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
1. Templates:
    * https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
    * https://rockfloat.com/post/learning-golang-templates.html
