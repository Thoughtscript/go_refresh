# experiments

> **README updated** 6.10.2020

> **NOTE:** Do not use white-spaces in file directory paths.

> **NOTE:** No configuration should be necessary to execute the samples below.

Supplied examples:

1. [1-asyncGo](./1-asyncGo)
    * Execute `$ go run asyncGo.go` by command-line after navigating to `1-asyncGo/src`
    * Can be executed without configuration
1. [2-httpServer](./2-httpServer)
    * Execute `$ go run httpServer.go` by command-line after navigating to `2-httpServer`
    * Can be executed without configuration
    * ```
      localhost:8888/print <- simple String to HTML writeout
      localhost:8888/write <- simple Model String to HTML writeout
      localhost:8888/public/index.html or localhost:8888/public/
      ```
1. [3-templateServer](./3-templateServer)
    * Execute `$ go run templateServer.go` by command-line after navigating to `4-templateServer`
    * Can be executed without configuration
    * Hard-coded `html` static file location
    * Files served by endpoint must be served through static server `public` path
    * ```
      localhost:7777/home <- simple String to HTML
      ```
1. [4-newMuxExample](./4-newMuxExample)
    * This replaces a previous version using new go features like `go get`
    * Execute `go get -u github.com/gorilla/mux` to download and install the dependencies
    * The default GOPATH on macOS will be `GOPATH="/Users/USER_NAME/go"` - packages installed via `go get` will appear here
    * Can be executed without configuration
    * ```
      localhost:3000/getall <- return JSON data
      ```
      
## Fantastic Resources

1. REST: http://www.blog.labouardy.com/build-restful-api-in-go-and-mongodb/
    * Gin: https://gin-gonic.github.io/gin/
1. Go Routines: 
    * https://medium.com/@nikolay.bystritskiy/how-i-tried-to-do-things-asynchronously-in-golang-40e0c1a06a66
    * https://gobyexample.com/goroutines
1. Static: http://www.alexedwards.net/blog/serving-static-sites-with-go
    * https://rockfloat.com/post/learning-golang-templates.html
    * https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
1. Templates:
    * https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
    * https://rockfloat.com/post/learning-golang-templates.html
1. Misc: https://blog.rubylearning.com/