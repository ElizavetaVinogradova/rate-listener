package service

import log "github.com/sirupsen/logrus"

type Tick struct {
	Timestamp int64
	Symbol    string
	Best_bid  float64
	Best_ask  float64
}

type RatesProvider interface {
	GetTicksBatch(batchSize int) ([]Tick, error)
}

type RatesRepository interface {
	CreateBatch(ticks []Tick) error
}

type TickService struct {
	ratesRepository RatesRepository
	ratesProvider   RatesProvider
	batchSize       int
}

func NewTickService(ratesRepository RatesRepository, ratesProvider RatesProvider, batchSize int) *TickService {
	return &TickService{
		ratesRepository: ratesRepository,
		ratesProvider:   ratesProvider,
		batchSize:       batchSize,
	}
}

func (s *TickService) Run() {
	for {
		ticks, err := s.ratesProvider.GetTicksBatch(s.batchSize)
		if err != nil {
			log.Errorf("Couldn't get ticks from provider: %s", err)
			continue
		}

		err = s.ratesRepository.CreateBatch(ticks)
		if err != nil {
			log.Errorf("Couldn't creare batch: %s", err)
		}
	}
}
