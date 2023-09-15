package kafka

import (
	"context"
	"rates-listener/internal/service"

	kafka "github.com/segmentio/kafka-go"
)

type BrokerWriter struct {
	writer *kafka.Writer
}

func NewBrokerWtiter(brokerAddresses []string, topic string) *BrokerWriter {
	config := kafka.WriterConfig{
		Brokers: brokerAddresses,
		Topic:   topic,
	}

	writer := kafka.NewWriter(config)
	return &BrokerWriter{writer: writer}
}

func (b *BrokerWriter) WriteBatch(ticks []service.Tick) error {
	byteKey := []byte("crypto-rate-key")
	byteValue, err := MarshalMessage(mapTickSliceToKafkaDTOSlice(ticks))
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   byteKey,
		Value: byteValue,
	}

	err = b.writer.WriteMessages(context.Background(), message)
	if err != nil {
		//todo move to service log.Fatal("Error writing to Kafka:", err)
		return err
	}
	return nil
}

// func mapTickSliceToKafkaDTOSlice(ticks []service.Tick) []TickKafkaDTO { //todo to add the initial capacity for the slice
// 	var kafkaTicks []TickKafkaDTO
// 	for _, tick := range ticks {
// 		kafkaTicks = append(kafkaTicks, mapTickToKafkaDTO(tick))
// 	}
// 	return kafkaTicks
// }

// func mapTickToKafkaDTO(tick service.Tick) TickKafkaDTO {
// 	var kafkaTick TickKafkaDTO
// 	kafkaTick.timestamp = tick.Timestamp
// 	kafkaTick.symbol = tick.Symbol
// 	kafkaTick.best_bid = tick.Best_bid
// 	kafkaTick.best_ask = tick.Best_ask
// 	return kafkaTick
// }
