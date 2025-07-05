# Go Practice: Server Programming

This repo is for practicing **Go (Golang)** with a focus on **backend and server-side programming**. It covers basic concepts and builds up to writing production-ready servers.

## 🚀 Why Go?

* Fast compilation and execution
* Built-in concurrency (goroutines, channels)
* Standard library support for **HTTP, WebSocket, and TCP servers**
* Simple, readable syntax

## 📋 Topics Covered

1. **Go Basics**

   * Variables, functions
   * Structs and interfaces
   * Error handling
   * Pointers
   * Packages & modules

2. **Networking & Server Programming**

   * HTTP servers (using `net/http`)
   * WebSocket servers (using `gorilla/websocket`)
   * TCP & UDP sockets
   * REST APIs
   * Middleware basics

3. **Concurrency**

   * Goroutines
   * Channels
   * Worker pools

4. **Project Ideas**

   * Simple HTTP chat server
   * WebSocket chat server
   * JSON-based REST API
   * File server
   * Health-check and monitoring service

## 📦 Key Libraries

* Standard library:

  * `net/http` (HTTP servers)
  * `net` (TCP/UDP)
* Popular external:

  * `github.com/gorilla/websocket` (WebSocket support)
  * `github.com/gin-gonic/gin` (HTTP framework, optional)

## ✅ How to Run

1️⃣ Install Go:
👉 [https://go.dev/dl/](https://go.dev/dl/)

2️⃣ Run any program:

```bash
go run filename.go
```

3️⃣ For modules:

```bash
go mod init mymodule
go get [package]
```


## Commands -- 

```bash
go mod init <file_name>  # to create the file name mod file that has the reqs 

go run .   #to run the file one file that has main fun 

go help   #go get the help for the function 

go mod tidy   #to fix the go.mod file 

go build  # to build the executible 

go test # to run the test in the code

go get .  #install the module from the go package

```

## 💡 First Example: HTTP Server

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
```

Links -

- https://go.dev/doc/tutorial/call-module-code

- https://pkg.go.dev/github.com/gorilla/websocket#section-readme

- https://gowebexamples.com/



