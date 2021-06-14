package main

import (
	"encoding/json"
	"news_portal/internal"
	processor "news_portal/internal/app/news_processor"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

func main() {
	logger := internal.InitLogger()
	zap.ReplaceGlobals(logger)

	kafkaConn := "localhost:9092"
	kafkaBrokers := []string{kafkaConn}
	master, err := internal.NewConsumer(kafkaBrokers)
	if err != nil {
		logger.Error("kafka consumer connection error", zap.Error(err))
	}
	producer, err := internal.NewProducer(kafkaBrokers)

	if err != nil {
		logger.Error("kafka producer connection error", zap.Error(err))
	}

	topic := "raw_news"

	//partions, err := master.Partitions(topic)
	if err != nil {
		logger.Error("kafka partions error", zap.Error(err))
	}
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		logger.Error("consumer error", zap.Error(err))
	}

	for i := range consumer.Messages() {
		//fmt.Println(string(i.Value))
		var page internal.RawPage
		err = json.Unmarshal(i.Value, &page)
		if err != nil {
			logger.Error("unmarshalling error", zap.Error(err))
		}
		//fmt.Println(page.URL)
		news, topic := processor.ProcessNews(&page)

		internal.ExportPage(producer, topic, news)
	}

}
