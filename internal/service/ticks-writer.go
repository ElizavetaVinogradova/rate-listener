package service

import log "github.com/sirupsen/logrus"

type RatesProvider interface {
	GetTicksBatch(batchSize int) ([]Tick, error)
}

type RatesBrokerWriter interface {
	WriteBatch(ticks []Tick) error
}

type TickWriterService struct {
	ratesProvider     RatesProvider
	ratesBrokerWriter RatesBrokerWriter
	batchSize         int
}

func NewTickWriterService(ratesProvider RatesProvider, ratesBrokerWriter RatesBrokerWriter, batchSize int) *TickWriterService {
	return &TickWriterService{
		ratesProvider:     ratesProvider,
		ratesBrokerWriter: ratesBrokerWriter,
		batchSize:         batchSize,
	}
}

func (s *TickWriterService) RunToKafka() {
	for {
		ticks, err := s.ratesProvider.GetTicksBatch(s.batchSize)
		if err != nil {
			log.Errorf("Couldn't get ticks from provider: %s", err)
			continue
		}

		err = s.ratesBrokerWriter.WriteBatch(ticks)
		if err != nil {
			log.Errorf("Error writing to Kafka: %s", err)
		}
	}
}
