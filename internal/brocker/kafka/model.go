package kafka

import (
	"encoding/json"
)

type TickKafkaDTO struct {
	Timestamp int64
	Symbol    string
	BestBid   float64
	BestAsk   float64
}

func MarshalMessage(ticks []TickKafkaDTO) ([]byte, error) {
	message, err := json.Marshal(&ticks)
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
