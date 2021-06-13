package internal

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type DB struct {
	DB *sqlx.DB
}

func NewDB() *DB {

	dbConn := "postgresql://user:pass@0.0.0.0:5432/postgres?sslmode=disable"
	db := sqlx.MustOpen("postgres", dbConn)

	err := db.Ping()
	if err != nil {
		zap.L().Error("db ping error", zap.Error(err))
	}

	return &DB{DB: db}
}

func InsertNews(db *sqlx.DB, news *NewsPage, topic string) {
	tables := GetTables(topic)
	for _, table := range tables {

		query := `INSERT INTO ` + table + ` (title, text, url) VALUES ($1, $2, $3);`

		_, err := db.Exec(query, news.Title, news.Text, news.Url)

		if err != nil {
			zap.L().Error("insert news error", zap.Error(err))
		} else {
			zap.L().Info("inserting news to " + table)
		}
	}

}

func GetTables(topic string) []string {
	switch topic {
	case "economy":
		return []string{"news_economy",
			"news_economy_politics",
			"news_science_economy",
			"news_economy_politics_science"}
	case "science":
		return []string{"news_science",
			"news_science_economy",
			"news_science_politics",
			"news_economy_politics_science"}
	case "politics":
		return []string{"news_politics",
			"news_science_politics",
			"news_economy_politics",
			"news_economy_politics_science"}

	}
	return []string{}
}
