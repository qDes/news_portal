package main

import (
	"encoding/json"
	"fmt"
	"news_portal/internal"
	"strconv"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type WebPage struct {
	HTML string
	URL  string
}

func main() {

	logger := internal.InitLogger()
	zap.ReplaceGlobals(logger)

	kafka_conn := "localhost:9092"
	kafkaBrokers := []string{kafka_conn}
	producer, err := NewProducer(kafkaBrokers)

	topic := "raw_news"
	if err != nil {
		logger.Error("kafka connection error", zap.Error(err))
	}

	var (
		page WebPage
		msg  *sarama.ProducerMessage
	)
	for i := 0; i <= 10000; i++ {
		page.URL = "sasi.ru/" + strconv.Itoa(i)
		page.HTML = "Text " + strconv.Itoa(i)
		msg = prepareMessage(topic, &page)
		fmt.Println(msg)
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			logger.Error("Kafka sending message error", zap.Error(err))
		} else {
			logger.Info(fmt.Sprintf("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset))
		}
	}

}

func NewProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func prepareMessage(topic string, record *WebPage) *sarama.ProducerMessage {
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
