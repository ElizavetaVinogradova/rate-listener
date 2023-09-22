package main

import (
	"fmt"
	"rates-listener/internal/brocker/kafka"
	"rates-listener/internal/repo/mysql"
	"rates-listener/internal/service"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		panic(fmt.Sprintf("Couldnt initialize the config file: %s", err))
	}
	setupLogging()

	broker := kafka.NewBrokerReader([]string{"localhost:9093"}, "ticks")

	repository, err := mysql.NewTickRepository(buildMySqlConfig())

	if err != nil {
		panic(fmt.Sprintf("Couldnt create Repository: %s", err))
	}

	viper.SetDefault("service.batchSize", 1)
	batchSize := viper.GetInt("service.batchSize")
	service.NewTickReaderService(broker, repository, batchSize).RunFromKafka()
}

func buildMySqlConfig() mysql.Config {
	return mysql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
}

func initConfig() error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
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
