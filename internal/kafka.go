package internal

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type RawPage struct {
	HTML     string
	URL      string
	IDSource int
	DTTM     string
}

func NewConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.ClientID ="news-processor"
	//config.Consumer.Return.Errors = true
	// Defaults to OffsetNewest.

	consumer, err := sarama.NewConsumer(brokers, config)

	return consumer, err
}

func NewProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func PrepareMessage(topic string, record *RawPage) *sarama.ProducerMessage {
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
