package apiserver

import (
	log "github.com/sirupsen/logrus"
)

type ApiServer struct {
	config Config
}

type Config struct {
	BindAddress string
}

func NewApiServer(config Config) *ApiServer {
	return &ApiServer{
		config: config,
	}
}

func (s *ApiServer) Start() error {
	log.Info("Start server")
	return nil
}
