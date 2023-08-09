package main

import (
	"encoding/json"
	"fmt"
	"log"
	"rates-listener/internal/client/coinbase"
)

func main() {
	url := "wss://ws-feed.exchange.coinbase.com"
	client, _ := coinbase.NewCoinBaseClient(url)
	defer client.Conn.Close()

	for {
		messageType, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("Received message type %d: %s", messageType, message)

		var tickDTO coinbase.TickClientDTO
		err2 := json.Unmarshal(message, &tickDTO)
		if err2 != nil {
			fmt.Println("Unmarshalling failed:", err)
			return
		}

	}
}
