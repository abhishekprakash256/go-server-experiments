/*
The function to make the websocket server 
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

func HomePage(w http.ResponseWriter , r *http.Request) {
	 fmt.Fprintf(w, "Home Page")

}


func wsEndpoint(w http.ResponseWriter, r *http.Request) {

    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    // upgrade this connection to a WebSocket
    // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {

		log.Println("Upgrade error:", err)
		return  // prevent using nil ws
		
	}
    // helpful log statement to show connections
    log.Println("Client Connected")

    reader(ws)
}

func reader(conn *websocket.Conn) {
    for {
    // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
    // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}


func setupRoutes() {

	http.HandleFunc("/", HomePage)
    http.HandleFunc("/ws", wsEndpoint)
}

func main() {

	fmt.Println("In the main")

	setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
