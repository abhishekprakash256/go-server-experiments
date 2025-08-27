/*

The server for one to one chat system 
using json 
*/



package main


import (

	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

)



// --- message format ----- 

type Message struct {

	From string `json:"from"`
	To string 	`json:"to"`
	Content string `json:"content"`
}


// ---- Global state ----- 

var clients = make(map[string]*websocket.Conn) 
var messages = make(chan Message)


// make the upgrader
var upgrader = websocket.Upgrader {

	ReadBufferSize : 1024 ,
	WriteBufferSize : 1024 , 
	CheckOrigin : func(r *http.Request) bool { return true } , 

}


// --- websocket handler ---- 

func wsEndpoint(w http.ResponseWriter , r *http.Request) {

	// Upgrade HTTP -> WS

	conn , err := upgrader.Upgrade(w, r, nil) 
	if err != nil {
		
		log.Println(err)
		return
	}


	// First message 
	_ , msgBytes , err := conn.ReadMessage()
	if err != nil {

		log.Println("Error reading ID: " , err )
		conn.Close()
		return 
	}

	var msg Message 

	if err := json.Unmarshal(msgBytes , &msg) ; err != nil {

		log.Println("Invalid ID Message" , err)
		conn.Close()
		return
	}

	clientID := msg.From

	clients[clientID] = conn

	log.Printf(" %s conncted \n" , clientID)


	// start listening to the message 
	go handleClinet(conn, clientID)

}



// ---- Handle Client message ---- 

func handleClinet(conn *websocket.Conn , clientID string )  {

	// runs after the function is running 
	defer func() {

		conn.Close()
		delete(clients , clientID)
		log.Println(" %s disconnedted\n " , clientID)

	}()

	for {
		
		_ , msgBytes , err := conn.ReadMessage()

		if err != nil {

			log.Println("Read error : ", err )

			break
		}


		var msg Message 

		if err := json.Unmarshal(msgBytes , &msg) ; err != nil {

			log.Println("Invalid Message " , err )

			continue
		}

		// enques the message 

		messages <- msg

	}



}


// ---- Message dispatcher ---- 
func handleMessages() {

	for {

		msg := <- messages

		if conn, ok := clients[msg.To]; ok {

			if err := conn.WriteJSON(msg); err != nil {

				log.Println("Send Error" , err)
				conn.Close()
				delete(clients , msg.To)

			}

		} else {
				log.Printf("User %s not connected \n" , msg.To)
			}

		}
}


func setupRoutes() {

	http.HandleFunc("/ws" , wsEndpoint)

}


func main() {

	fmt.Println("Server running on :8080")

	setupRoutes()

	go handleMessages()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

