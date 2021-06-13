package main

import (
	"fmt"
	"news_portal/internal"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func main() {

	logger := internal.InitLogger()
	zap.ReplaceGlobals(logger)

	conn := "postgresql://user:pass@0.0.0.0:5432/postgres?sslmode=disable"
	db := sqlx.MustOpen("postgres", conn)

	res := internal.GetFeedByTopic(db, "science")
	fmt.Println(res)
}
