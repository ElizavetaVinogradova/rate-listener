package apiserver

import (
	"fmt"
	"net/http"
	"rates-listener/internal/service"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type ApiServer struct {
	tickService service.TickService
	config      Config
}

type Config struct {
	BindAddress string
}

func NewApiServer(config Config, tickService service.TickService) *ApiServer {
	return &ApiServer{
		config:      config,
		tickService: tickService,
	}
}

func (s *ApiServer) Start() error {
	log.Info("Start server")
	mux := http.NewServeMux()
	mux.HandleFunc("/ticks", s.getTickById)
	err := http.ListenAndServe(s.config.BindAddress, mux)
	return err
}

func (s *ApiServer) getTickById(w http.ResponseWriter, r *http.Request) {
	log.Info("got /ticks request")

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Errorf("Missing id parameter, %d", http.StatusBadRequest)
		return
	}

	idValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Errorf("id parsing, %d", http.StatusInternalServerError)
		return
	}

	tick, err := s.tickService.GetTickById(idValue)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting tick by id: %s", err), http.StatusInternalServerError)
		return
	}
	tickAPI := mapTickToTickAPI(tick)
	jsonResponse, err := MarshalRequest(tickAPI)
	if err != nil {
		http.Error(w, "Error marshalling", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func mapTickToTickAPI(tick service.Tick) TickAPI {
	var tickAPI TickAPI
	tickAPI.Timestamp = tick.Timestamp
	tickAPI.Symbol = tick.Symbol
	tickAPI.BestBid = tick.BestBid
	tickAPI.BestAsk = tick.BestAsk
	return tickAPI
}
