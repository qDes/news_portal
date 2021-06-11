package main

import (
	"news_portal/internal"
	"news_portal/internal/app/scraper"
	"sync"

	"go.uber.org/zap"
)

func main() {

	logger := internal.InitLogger()
	zap.ReplaceGlobals(logger)

	var (
		wg sync.WaitGroup
		//counter atomic.Int32

	)

	chOut := make(chan *internal.RawPage, 10)

	wg.Add(1)
	go scraper.PostProcessLoop(chOut)
	for _, url := range scraper.GetScraperConfig() {
		chUrls := make(chan *internal.RawPage, 10)
		go scraper.ScanLoop(url, chUrls)
		go scraper.ProcessLoop(chUrls, chOut)
	}
	wg.Wait()

	// kafka test
	/*
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

	*/

}
