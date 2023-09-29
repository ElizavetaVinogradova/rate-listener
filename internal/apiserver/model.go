package apiserver

import (
	"encoding/json"
	"fmt"
)

type TickAPI struct {
	Id        int64   `json:"id"`
	Timestamp int64   `json:"timestamp"`
	Symbol    string  `json:"symbol"`
	BestBid   float64 `json:"best_bid"`
	BestAsk   float64 `json:"best_ask"`
}

func MarshalRequest(request TickAPI) ([]byte, error) {
	requestMsg, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshall request to API: %w", err)
	}
	return requestMsg, nil
}

func UnmarshalResponse(message []byte) (TickAPI, error) {
	var tickAPI TickAPI
	err := json.Unmarshal(message, &tickAPI)
	if err != nil {
		return TickAPI{}, fmt.Errorf("unmarshall request to coinbase: %w", err)
	}
	return tickAPI, nil
}
