package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	url := "wss://ws-feed.exchange.coinbase.com"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	subscribeMsg := []byte(`{
        "type": "subscribe",
        "channels": [
            {
                "name": "ticker",
                "product_ids": ["BTC-USD"]
            }
        ]
    }`)
	err = conn.WriteMessage(websocket.TextMessage, subscribeMsg)
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("Received message type %d: %s", messageType, message)
	}
}
