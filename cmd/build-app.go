package cmd

import (
	"rates-listener/internal/apiserver"
	"rates-listener/internal/repo/mysql"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}

func SetupLogging() {
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

func BuildMySqlConfig() mysql.Config {
	return mysql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
}

func BuildApiServerConfig() apiserver.Config {
	return apiserver.Config{
		BindAddress: viper.GetString("apiserver.bindAddress"),
	}
}
