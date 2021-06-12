package main

import (
	"news_portal/internal"
	"news_portal/internal/app/scraper"
	"sync"
	"time"

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
		wg sync.WaitGroup
	)

	chOut := make(chan *internal.RawPage, 100)

	wg.Add(1)
	go scraper.PostProcessLoop(chOut, producer, topic)
	for _, url := range scraper.GetScraperConfig() {
		chUrls := make(chan *internal.RawPage, 100)
		go scraper.ScanLoop(url, chUrls)
		go scraper.ProcessLoop(chUrls, chOut)
		time.Sleep(time.Duration(10) * time.Second)
	}
	wg.Wait()


}
