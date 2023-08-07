package service

import "time"

type Tick struct {
	timestamp time.Time
	symbol    string
	best_bid  float64
	best_ask  float64
}

type TickService interface {
	Get() Tick
	Post() Tick
	Put() Tick
	Delete() Tick
}
