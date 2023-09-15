package kafka

import (
	"encoding/json"
)

type TickKafkaDTO struct {
	timestamp int64
	symbol    string
	best_bid  float64
	best_ask  float64
}

func MarshalMessage(ticks []TickKafkaDTO) ([]byte, error) {
	message, err := json.Marshal(ticks)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func UnmarshalMessage(messages []byte) ([]TickKafkaDTO, error) {
	var ticks []TickKafkaDTO
	err := json.Unmarshal(messages, &ticks)
	if err != nil {
		return nil, err
	}
	return ticks, nil
}
