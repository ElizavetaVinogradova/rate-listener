package main

import (
	"fmt"
	"rates-listener/cmd"
	"rates-listener/internal/brocker/kafka"
	"rates-listener/internal/client/coinbase"
	"rates-listener/internal/service"

	"github.com/spf13/viper"
)

func main() {
	if err := cmd.InitConfig(); err != nil {
		panic(fmt.Sprintf("Couldnt initialize the config file: %s", err))
	}
	cmd.SetupLogging()

	url := viper.GetString("coinBaseClient.url")
	client, err := coinbase.NewCoinBaseClient(url)
	if err != nil {
		panic(fmt.Sprintf("Couldnt create CoinBase Client: %s", err))
	}
	defer client.Conn.Close()

	broker := kafka.NewBrokerWriter([]string{"localhost:9093"}, "ticks")

	viper.SetDefault("service.batchSize", 1)
	batchSize := viper.GetInt("service.batchSize")
	service.NewTickWriterService(client, broker, batchSize).RunToKafka()
}
