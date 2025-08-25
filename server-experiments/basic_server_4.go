/*
The server for connecting the clinets 
*/

package main 


import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {

	ReadBufferSize : 1024 ,
	WriteBufferSize : 1024 , 
	CheckOrigin : func(r *http.Request) bool { return true } , 

}

// Connnected clients
var clients  = make(map[*websocket.Conn]bool)

//broadcast channel 
var broadcast = make(chan []byte)

func main() {

	http.HandleFunc("/ws" , handleConnections)

	go handleMessages()

	fmt.Println("Server started on : 8080 ")

	err := http.ListenAndServe(": 8080" , nil)

	if err != nil {

		log.Fatal("ListenandServer" , err)
	}

}


func handleConnections(w http.ResponseWriter , r * http.Request) {

	// Upgrade initial GET 
	ws , err := upgrader.Upgrade(w,r, nil)

	if err != nil {

		log.Println("Upgrade error: " , err)

		return
	}

	defer ws.Close()
	
	// Register new clinet
	clients[ws] = true
	log.Println("New Clinet connected ")


	for {
		// Read the message from the client

		_, msg , err := ws.ReadMessage()

		if err != nil {

			log.Println("Read Error")
		}

		broadcast <- msg
	}

}


func handleMessages() {

	for {

		// Grab next message from broadcast channel 

		msg := <- broadcast

		//send to all connected clients

		for client := range clients {

			err := client.WriteMessage( websocket.TextMessage, msg)

			if err != nil {
				
				log.Println("write error : ", err)

				client.Close()

				delete(clients,client)
			}
		}


	}


}

