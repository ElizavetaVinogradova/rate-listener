package kafka

import (
	"encoding/json"
	"fmt"
)

type TickKafkaDTO struct {
	Id        int64   `json:"id"`
	Timestamp int64   `json:"timestamp"`
	Symbol    string  `json:"symbol"`
	BestBid   float64 `json:"best_bid"`
	BestAsk   float64 `json:"best_ask"`
}

func MarshalMessage(ticks []TickKafkaDTO) ([]byte, error) {
	message, err := json.Marshal(&ticks)
	if err != nil {
		return nil, fmt.Errorf("marshall message to Kafka: %w", err)
	}
	return message, nil
}

func UnmarshalMessage(messages []byte) ([]TickKafkaDTO, error) {
	var ticks []TickKafkaDTO
	err := json.Unmarshal(messages, &ticks)
	if err != nil {
		return nil, fmt.Errorf("unmarshall message to Kafka: %w", err)
	}
	return ticks, nil
}
