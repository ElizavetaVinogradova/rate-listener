package main

import (
	"fmt"
	"rates-listener/cmd"
	"rates-listener/internal/brocker/kafka"
	"rates-listener/internal/repo/mysql"
	"rates-listener/internal/service"

	"github.com/spf13/viper"
)

func main() {
	if err := cmd.InitConfig(); err != nil {
		panic(fmt.Sprintf("Couldnt initialize the config file: %s", err))
	}
	cmd.SetupLogging()

	broker := kafka.NewBrokerReader([]string{"localhost:9093"}, "ticks")
	defer broker.Close()

	repository, err := mysql.NewTickRepository(cmd.BuildMySqlConfig())
	if err != nil {
		panic(fmt.Sprintf("Couldnt create Repository: %s", err))
	}
	defer repository.Close()

	viper.SetDefault("service.batchSize", 1)
	batchSize := viper.GetInt("service.batchSize")
	service.NewTickReaderService(broker, repository, batchSize).RunReadingFromBroker()
}
