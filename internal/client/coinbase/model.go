package coinbase

import (
	"encoding/json"
	"fmt"
	"rates-listener/internal/service"
	"strconv"
	"time"
)

type TickClientDTO struct {
	Type        string    `json:"type"`
	Sequence    int64     `json:"sequence"`
	ProductID   string    `json:"product_id"`
	Price       string    `json:"price"`
	Open24H     string    `json:"ope n_24h"`
	Volume24H   string    `json:"volume_24h"`
	Low24H      string    `json:"low_24h"`
	High24H     string    `json:"high_24h"`
	Volume30D   string    `json:"volume_30d"`
	BestBid     string    `json:"best_bid"`
	BestBidSize string    `json:"best_bid_size"`
	BestAsk     string    `json:"best_ask"`
	BestAskSize string    `json:"best_ask_size"`
	Side        string    `json:"side"`
	Time        time.Time `json:"time"`
	TradeID     int64     `json:"trade_id"`
	LastSize    string    `json:"last_size"`
}

func (dto *TickClientDTO) toServiceTick() service.Tick {
	var tick service.Tick
	tick.Timestamp = dto.Time.Unix()
	tick.Symbol = dto.ProductID
	tick.BestBid, _ = strconv.ParseFloat(dto.BestBid, 64)
	tick.BestAsk, _ = strconv.ParseFloat(dto.BestAsk, 64)
	return tick
}

type RequestMessage struct {
	Type     string    `json:"type"`
	Channels []Channel `json:"channels"`
}

type Channel struct {
	Name       string   `json:"name"`
	ProductIDS []string `json:"product_ids"`
}

func MarshalRequest(request RequestMessage) ([]byte, error) {
	requestMsg, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshall request to coinbase: %w", err)
	}
	return requestMsg, nil
}

func UnmarshalResponse(message []byte) (TickClientDTO, error) {
	var tickDTO TickClientDTO
	err := json.Unmarshal(message, &tickDTO)
	if err != nil {
		return TickClientDTO{}, fmt.Errorf("unmarshall request to coinbase: %w", err)
	}
	return tickDTO, nil
}
