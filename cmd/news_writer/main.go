package main

import (
	"encoding/json"
	"fmt"
	"news_portal/internal"
	"sync"

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

	topicEconomy := "economy"
	topicPolitics := "politics"
	topicScience := "science"

	partions, err := master.Partitions(topicEconomy)
	if err != nil {
		logger.Error("kafka partions error", zap.Error(err))
	}

	consumerEc := GetConsumer(master, topicEconomy, partions)
	consumerPo := GetConsumer(master, topicPolitics, partions)
	consumerSc := GetConsumer(master, topicScience, partions)

	var wg sync.WaitGroup
	wg.Add(1)
	go ReadTopic(consumerEc)
	go ReadTopic(consumerPo)
	go ReadTopic(consumerSc)
	wg.Wait()

}

func GetConsumer(master sarama.Consumer, topic string, partions []int32) sarama.PartitionConsumer {
	consumer, err := master.ConsumePartition(topic, partions[0], sarama.OffsetOldest)
	if err != nil {
		zap.L().Error("consumer economy error", zap.Error(err))
	}
	return consumer
}

func ReadTopic(consumer sarama.PartitionConsumer) {
	for i := range consumer.Messages() {
		var page internal.NewsPage
		err := json.Unmarshal(i.Value, &page)
		if err != nil {
			zap.L().Error("unmarshalling error", zap.Error(err))
		}
		fmt.Println(page.Title)
	}
}