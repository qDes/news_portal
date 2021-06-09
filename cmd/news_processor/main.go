package main

import (
	"news_portal/internal"

	"go.uber.org/zap"
)

func main() {
	logger := internal.InitLogger()
	zap.ReplaceGlobals(logger)


}

func NewConsumer() {}