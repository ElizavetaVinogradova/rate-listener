package main

import (
	"fmt"
	"rates-listener/internal/brocker/kafka"
	"rates-listener/internal/client/coinbase"
	"rates-listener/internal/service"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		panic(fmt.Sprintf("Couldnt initialize the config file: %s", err))
	}
	setupLogging()

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

func initConfig() error {
	//viper.AddConfigPath("configs")
	// viper.AddConfigPath("./configs")
	// viper.SetConfigName("config")
	// viper.SetConfigType("yml")
    viper.SetConfigFile("C:\\Users\\liza\\GoProjects\\rates-listener\\configs\\config.yml")
	return viper.ReadInConfig()
}

func setupLogging() {
	logLevel := viper.GetString("logLevel")

	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
