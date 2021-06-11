package main

import (
	"fmt"
	"news_portal/internal"
	"strconv"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)


func main() {

	logger := internal.InitLogger()
	zap.ReplaceGlobals(logger)

	kafkaConn := "localhost:9092"
	kafkaBrokers := []string{kafkaConn}
	producer, err := internal.NewProducer(kafkaBrokers)

	topic := "raw_news"
	if err != nil {
		logger.Error("kafka connection error", zap.Error(err))
	}

	var (
		page internal.RawPage
		msg  *sarama.ProducerMessage
	)
	for i := 0; i <= 10000; i++ {
		page.URL = "sasi.ru/" + strconv.Itoa(i)
		page.HTML = "Text " + strconv.Itoa(i)
		msg = internal.PrepareMessage(topic, &page)
		fmt.Println(msg)
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			logger.Error("Kafka sending message error", zap.Error(err))
		} else {
			logger.Info(fmt.Sprintf("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset))
		}
	}

}


