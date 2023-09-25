package service

import (
	log "github.com/sirupsen/logrus"
)

type RatesBrokerReader interface {
	ReadBatch() ([]Tick, error)
}

type RatesRepository interface {
	CreateBatch(ticks []Tick) error
}

type TickReaderService struct {
	ratesBrokerReader RatesBrokerReader
	ratesRepository   RatesRepository
	batchSize         int
}

func NewTickReaderService(ratesBrokerReader RatesBrokerReader, ratesRepository RatesRepository, batchSize int) *TickReaderService {
	return &TickReaderService{
		ratesBrokerReader: ratesBrokerReader,
		ratesRepository:   ratesRepository,
		batchSize:         batchSize,
	}
}

func (s *TickReaderService) RunReadingFromBroker() {
	for {
		ticks, err := s.ratesBrokerReader.ReadBatch()
		if err != nil {
			log.Errorf("Couldn't get ticks from Kafka: %s", err)
			continue
		}

		err = s.ratesRepository.CreateBatch(ticks)
		if err != nil {
			log.Errorf("Couldn't create batch: %s", err)
		}
	}
}
