package service

type RatesRepository interface {
	CreateBatch(ticks []Tick) error
	GetTickById(id int64) (Tick, error)
}

type Tick struct {
	Id        int64
	Timestamp int64
	Symbol    string
	BestBid   float64
	BestAsk   float64
}
