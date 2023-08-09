package main

import (
	"rates-listener/internal/client/coinbase"
)

func main() {
	url := "wss://ws-feed.exchange.coinbase.com"
	client, _ := coinbase.NewCoinBaseClient(url)
	defer client.Conn.Close()

	const batchSize = 10
	messages := make([]coinbase.TickClientDTO, batchSize)

	for {
		coinbase.ReadBatchOfMessages(client, messages)
	}
}
