package kafka

import (
	"context"
	"fmt"
	"rates-listener/internal/service"

	log "github.com/sirupsen/logrus"

	kafka "github.com/segmentio/kafka-go"
)

type BrokerWriter struct {
	writer *kafka.Writer
}

func NewBrokerWriter(brokerAddresses []string, topic string) *BrokerWriter {
	config := kafka.WriterConfig{
		Brokers: brokerAddresses,
		Topic:   topic,
	}

	writer := kafka.NewWriter(config)
	return &BrokerWriter{writer: writer}
}

func (b *BrokerWriter) Close() {
	b.writer.Close()
}

func (b *BrokerWriter) WriteBatch(ticks []service.Tick) error {
	log.Debugf("Got ticks from coinbase: %s", fmt.Sprintf("%v", ticks))

	byteKey := []byte("crypto-rate-key")
	byteValue, err := MarshalMessage(mapTickSliceToKafkaDTOSlice(ticks))
	if err != nil {
		return err
	}
	log.Debugf("marshaled message for Kafka: %s", byteValue)

	message := kafka.Message{
		Key:   byteKey,
		Value: byteValue,
	}

	err = b.writer.WriteMessages(context.Background(), message)
	if err != nil {
		return fmt.Errorf("write message to Kafka: %w", err)
	}
	return nil
}

func mapTickSliceToKafkaDTOSlice(ticks []service.Tick) []TickKafkaDTO {
	kafkaTicks := make([]TickKafkaDTO, 0, len(ticks))
	for _, tick := range ticks {
		kafkaTicks = append(kafkaTicks, mapTickToKafkaDTO(tick))
	}
	return kafkaTicks
}

func mapTickToKafkaDTO(tick service.Tick) TickKafkaDTO {
	var kafkaTick TickKafkaDTO
	kafkaTick.Timestamp = tick.Timestamp
	kafkaTick.Symbol = tick.Symbol
	kafkaTick.BestBid = tick.BestBid
	kafkaTick.BestAsk = tick.BestAsk

	log.Debugf("Tick mapped to kafkaDTO: %s", fmt.Sprintf("%v", kafkaTick))
	return kafkaTick
}
