package internal

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type WebPage struct {
	HTML string
	URL  string
}

func NewProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func PrepareMessage(topic string, record *WebPage) *sarama.ProducerMessage {
	b, err := json.Marshal(record)
	if err != nil {
		zap.L().Error("prepareMessage error", zap.String("function", "prepareMessage"), zap.Error(err))
		return nil
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(string(b)),
	}

	return msg
}
