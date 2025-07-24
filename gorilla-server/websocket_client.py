"""
make the websocket client
"""

import websocket

def on_message(ws, message):
    print("Received from server:", message)

def on_error(ws, error):
    print("Error:", error)

def on_close(ws, close_status_code, close_msg):
    print("Connection closed")

def on_open(ws):
    print("Connected to server. Type messages and press Enter (Ctrl+C to quit).")
    def run():
        while True:
            msg = input("> ")  # take user input
            ws.send(msg)
    import threading
    threading.Thread(target=run).start()



# Create a WebSocket connection
ws = websocket.WebSocketApp(
    "ws://localhost:8080/ws",
    on_open=on_open,
    on_message=on_message,
    on_error=on_error,
    on_close=on_close
)

# Run it (this blocks)
ws.run_forever()


