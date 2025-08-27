/*
making the go client for connection with the 

*/

package main 

import (
	//"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"

)



// --- message format ----- 

type Message struct {

	From string `json:"from"`
	To string 	`json:"to"`
	Content string `json:"content"`
}



func main() {

	url := "ws://localhost:8080/ws"

	conn , _ , err := websocket.DefaultDialer.Dial(url , nil )

	if err != nil {
		
		log.Fatal("Dial error :" , err )
	}

	defer conn.Close()

	// ---identify the client---

	fmt.Print("Enter your ID: ")
	var myID string

	fmt.Scanln(&myID)

	identify := Message{From : myID , Content: "register" }

	conn.WriteJSON(identify)

	// --- Goroutine to read messages --- 

	go func() {

		for {


			var msg Message

			if err := conn.ReadJSON(&msg); err != nil {
				
				log.Println("Read error " , err)

				return

			}

			fmt.Println("\n%s → %s: %s\n", msg.From, msg.To, msg.Content)
		} 


	}()

	// ---- Handle -----

	interrupt := make(chan os.Signal, 1 )

	signal.Notify(interrupt , os.Interrupt)


	// ---- send loop ---- 

	for {

		fmt.Print("send (to msg) : ")

		var to, content string

		fmt.Scanln(&to , &content)

		msg := Message{From: myID, To: to, Content: content}

		err := conn.WriteJSON(msg)

		if err != nil {

			log.Println("Write error : " , err )
			break
		}

	}

}