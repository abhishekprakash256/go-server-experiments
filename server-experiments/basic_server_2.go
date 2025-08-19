/*
this is the websocket server in go 
*/

package main 


import (

	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	
)


var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}



func HomePage(w http.ResponseWriter , r * http.Request) {
	/*
	This function is making a home page
	*/
	fmt.Fprintf(w, "<h1>Home Page</h1>")
}



func wsEndpoint(w http.ResponseWriter , r * http.Request) {
	/*
	This function is websocket page
	*/
	
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	_, err := upgrader.Upgrade(w, r, nil)

	if err != nil {

		log.Println(err)

	}


}

func setupRoutes() {
	/*
	This function is setting up the routes
	*/

	fmt.Println("The server is started on localhost : 8080")

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	/*
	The main function and setting up the routes
	*/

	setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}





