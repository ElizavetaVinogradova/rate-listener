package kafka

import (
	"rates-listener/internal/service"

	kafka "github.com/segmentio/kafka-go"
)

type BrokerReader struct {
	reader *kafka.Reader
}

func NewBrokerReader(brokerAddresses []string, topic string) *BrokerReader {
	config := kafka.ReaderConfig{
		Brokers: brokerAddresses,
		Topic:   topic,
	}

	reader := kafka.NewReader(config)
	return &BrokerReader{reader: reader}
}

func (b *BrokerWriter) ReadBatch(ticks []service.Tick) error {
	return nil
}

func mapTickSliceToKafkaDTOSlice(ticks []service.Tick) []TickKafkaDTO { //todo to add the initial capacity for the slice
	var kafkaTicks []TickKafkaDTO
	for _, tick := range ticks {
		kafkaTicks = append(kafkaTicks, mapTickToKafkaDTO(tick))
	}
	return kafkaTicks
}

func mapTickToKafkaDTO(tick service.Tick) TickKafkaDTO {
	var kafkaTick TickKafkaDTO
	kafkaTick.timestamp = tick.Timestamp
	kafkaTick.symbol = tick.Symbol
	kafkaTick.best_bid = tick.Best_bid
	kafkaTick.best_ask = tick.Best_ask
	return kafkaTick
}
