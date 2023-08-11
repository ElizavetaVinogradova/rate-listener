package main

import (
	"fmt"
	"rates-listener/internal/client/coinbase"
	"rates-listener/internal/repo/mysql"
	"rates-listener/internal/service"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		panic(fmt.Sprintf("Couldnt initialize the config file: %s", err))
	}
	url := viper.GetString("service.url")
	client, err := coinbase.NewCoinBaseClient(url)
	if err != nil {
		panic(fmt.Sprintf("Couldnt create CoinBase Client: %s", err))
	}

	repository, err := mysql.NewTickRepository(mysql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		panic(fmt.Sprintf("Couldnt create Repository: %s", err))
	}
	defer client.Conn.Close()

	viper.SetDefault("service.batchSize", 1)
	batchSize := viper.GetInt("service.batchSize")
	service.NewTickService(repository, client, batchSize).Run()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
