package kafka

import (
	kafka "github.com/segmentio/kafka-go"
)

type BrokerWriter struct {
	kafka kafka.Writer
}

func NewBrokerWtiter(brokerAddresses []string, topic string) {
	config := kafka.WriterConfig{
		Brokers: brokerAddresses,
		Topic:   topic,
	}

	writer := kafka.NewWriter(config)
	return &BrokerWriter{writer: writer}
}

func WriteBatch(ticks []Tick) error {

}
