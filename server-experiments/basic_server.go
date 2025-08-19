/*
this is the basic server in go
*/

package main 

import (

	"fmt"
	"log"
	"net/http"
	
)


func HomePage(w http.ResponseWriter , r * http.Request) {
	/*
	This function is making a home page
	*/
	fmt.Fprintf(w, "<h1>Home Page</h1>")
}


func setupRoutes() {
	/*
	This function is setting up the routes
	*/

	fmt.Println("The server is started on localhost : 8080")

	http.HandleFunc("/", HomePage)
}



func main() {
	/*
	The main function and setting up the routes
	*/

	setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}