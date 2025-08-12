/*
this is a file to test the go gorilla server 

the ws is not a websocket handler and can't be used as one , need an upgrader


https://tutorialedge.net/golang/go-websocket-tutorial/


*/

package main 


import (
	"fmt"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter , r *http.Request) {
	 fmt.Fprintf(w, "Home Page")

}

func wsEndPoint(w http.ResponseWriter, r *http.Request ) {

	fmt.Fprintf(w, "Hello World")
}


func setupRoutes() {

	http.HandleFunc("/", HomePage)
    http.HandleFunc("/ws", wsEndPoint)
}

func main() {

	fmt.Println("In the main")

	setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}