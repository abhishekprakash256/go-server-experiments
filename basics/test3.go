package main

import (
    "fmt"
    "time"
)

func sayHello(done chan bool) {
    fmt.Println("Hello from goroutine!")
    time.Sleep(2 * time.Second)
    done <- true
}

func main() {
    done := make(chan bool)
    go sayHello(done) // Start goroutine
    <-done // Wait for goroutine to finish
    fmt.Println("Main function finishes.")
}
