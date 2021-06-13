package feedServer

import (
	feed "news_portal/api/proto"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type feedSrv struct {
	feed.UnimplementedFeedServiceServer
	db *sqlx.DB
}

func NewFeedSrv(conn string) *feedSrv {
	db := sqlx.MustOpen("postgres", conn)
	s := &feedSrv{db: db}
	return s
}