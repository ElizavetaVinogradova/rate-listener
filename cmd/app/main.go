package main

import (
	"fmt"
	"rates-listener/internal/client/coinbase"
	"rates-listener/internal/repo/mysql"
	"rates-listener/internal/service"
)

func main() {

	config := mysql.Config{
		Host:     "localhost",
		Port:     "3306",
		Username: "user",
		Password: "1234",
		DBName:   "ticks",
		SSLMode:  "disable",
	}

	url := "wss://ws-feed.exchange.coinbase.com"
	client, err := coinbase.NewCoinBaseClient(url)
	if err != nil {
		panic(fmt.Sprintf("Couldnt create CoinBase Client: %s", err))
	}
	repository, err := mysql.NewTickRepository(config)
	if err != nil {
		panic(fmt.Sprintf("Couldnt create Repository: %s", err))
	}
	defer client.Conn.Close()

	batchSize := 10
	service.NewTickService(repository, client, batchSize).Run()
}
