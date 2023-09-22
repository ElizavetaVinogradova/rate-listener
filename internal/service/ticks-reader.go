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

func NewTickReaderService(ratesBrokerReader RatesBrokerReader, ratesRepository RatesRepository, batchSize int) *TickReaderService { //тип возвращаемого значения - поинтер (заявление, что тип - поинтер)
	return &TickReaderService{ //взять поинтер на структуру (действие - взятие поинтера)
		ratesBrokerReader: ratesBrokerReader,
		ratesRepository:   ratesRepository,
		batchSize:         batchSize,
	}
}

//звездочка говорит о том, что метод вызывается на поинтерах на структуру. То есть, если там будет просто структура, то метод нельзя будет вызвать.
//написала метод для поинтера, потому что в го принято делать управляющие структуры поинтерами
func (s *TickReaderService) RunFromKafka() {
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
