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

type RatesBrokerWriter interface {
	WriteBatch(ticks []Tick) error
}

type TickService struct {
	ratesProvider     RatesProvider
	ratesBrokerWriter RatesBrokerWriter
	batchSize         int
}

func NewTickService(ratesProvider RatesProvider, ratesBrokerWriter RatesBrokerWriter, batchSize int) *TickService { //тип возвращаемого значения - поинтер (заявление, что тип - поинтер)
	return &TickService{ //взять поинтер на структуру (действие - взятие поинтера)
		ratesProvider:     ratesProvider,
		ratesBrokerWriter: ratesBrokerWriter,
		batchSize:         batchSize,
	}
}

func (s *TickService) RunToKafka() {
	for {
		ticks, err := s.ratesProvider.GetTicksBatch(s.batchSize)
		if err != nil {
			log.Errorf("Couldn't get ticks from provider: %s", err)
			continue
		}

		err = s.ratesBrokerWriter.WriteBatch(ticks)
		if err != nil {
			log.Errorf("Couldn't write batch: %s", err)
		}
	}
}
