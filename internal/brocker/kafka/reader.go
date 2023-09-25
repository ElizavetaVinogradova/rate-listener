package kafka

import (
	"context"
	"fmt"
	"rates-listener/internal/service"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
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

func (b *BrokerReader) Close(){
	b.reader.Close()
}

func (b *BrokerReader) ReadBatch() ([]service.Tick, error) {
	message, err := b.reader.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}

	kafkaDTOs, err := UnmarshalMessage(message.Value)
	if err != nil {
		return nil, err
	}
	log.Debugf("Ticks from Kafka Unmarshalled: %s", fmt.Sprintf("%v", kafkaDTOs))

	return mapKafkaDTOSliceToTicksDTO(kafkaDTOs), nil
}

func mapKafkaDTOSliceToTicksDTO(kafkaDTOs []TickKafkaDTO) []service.Tick {
	ticks := make([]service.Tick, 0, len(kafkaDTOs))
	for _, dto := range kafkaDTOs {
		ticks = append(ticks, mapKafkaDTOToTick(dto))
	}
	log.Debugf("Mapped KafkaDTO slice to Ticks slice: %s", fmt.Sprintf("%v", ticks))
	return ticks
}

func mapKafkaDTOToTick(dto TickKafkaDTO) service.Tick {
	var tick service.Tick
	tick.Timestamp = dto.Timestamp
	tick.Symbol = dto.Symbol
	tick.BestBid = dto.BestBid
	tick.BestAsk = dto.BestAsk
	return tick
}
