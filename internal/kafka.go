package internal

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type RawPage struct {
	HTML     string
	URL      string
	IDSource int
	DTTM     string
}

type NewsPage struct {
	Title string
	Text  string
	Url   string
	Dttm  string
}

func NewConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.ClientID = "news-processor"
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

func PrepareMessage(topic string, record interface{}) *sarama.ProducerMessage {
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

func ExportPage(producer sarama.SyncProducer, topic string, page interface{}) {
	msg := PrepareMessage(topic, page)

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		zap.L().Error("Kafka sending message error", zap.Error(err))
	} else {
		zap.L().Info(fmt.Sprintf("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset))
	}
}
