package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)



var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        // Allow all origins (dev only)
        return true
    },
}




func handler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		// Read message
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)

		// Echo message back
		err = conn.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}		
	}
}



func main() {
	http.HandleFunc("/ws", handler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
