package kafka

import (
	"encoding/json"
	"fmt"
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
