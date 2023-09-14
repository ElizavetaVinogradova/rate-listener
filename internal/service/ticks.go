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

type RatesBrokerWtiter interface {
	WriteBatch(ticks []Tick) error
}

type TickService struct {
	ratesBrokerWtiter RatesBrokerWtiter
	ratesRepository   RatesRepository
	ratesProvider     RatesProvider
	batchSize         int
}

func NewTickService(ratesBrokerWtiter RatesBrokerWtiter, ratesRepository RatesRepository, ratesProvider RatesProvider, batchSize int) *TickService { //тип возвращаемого значения - поинтер (заявление, что тип - поинтер)
	return &TickService{ //взять поинтер на структуру (действие - взятие поинтера)
		ratesBrokerWtiter: ratesBrokerWtiter,
		ratesRepository:   ratesRepository,
		ratesProvider:     ratesProvider,
		batchSize:         batchSize,
	}
}

//звездочка говорит о том, что метод вызывается на поинтерах на структуру. То есть, если там будет просто структура, то метод нельзя будет вызвать.
//написала метод для поинтера, потому что в го принято делать управляющие структуры поинтерами
// func (s *TickService) Run() {
// 	for {
// 		ticks, err := s.ratesProvider.GetTicksBatch(s.batchSize)
// 		if err != nil {
// 			log.Errorf("Couldn't get ticks from provider: %s", err)
// 			continue
// 		}

// 		err = s.ratesRepository.CreateBatch(ticks)
// 		if err != nil {
// 			log.Errorf("Couldn't create batch: %s", err)
// 		}
// 	}
// }

func (s *TickService) Run() {
	for {
		ticks, err := s.ratesProvider.GetTicksBatch(s.batchSize)
		if err != nil {
			log.Errorf("Couldn't get ticks from provider: %s", err)
			continue
		}

		err = s.ratesBrokerWtiter.WriteBatch(ticks)
		if err != nil {
			log.Errorf("Couldn't write batch: %s", err)
		}
	}
}
