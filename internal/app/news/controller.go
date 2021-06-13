package news

import (
	"context"
	"fmt"
	"net/http"
	feed "news_portal/api/proto"

	"go.uber.org/zap"
)

var (
	svc = GetSvc()
)
func Index(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello!"))
}

func Science(resp http.ResponseWriter, req *http.Request) {

	ctx := context.Background()
	res, err := svc.FeedClient.GetFeed(ctx, &feed.GetFeedRequest{Topic: "science"})
	if err != nil {
		zap.L().Error("science handler err on grpc call", zap.Error(err))
	}
	fmt.Println(res.Feed)
	resp.Write([]byte("Hello Science!"))
}

func Politics(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello Politics!"))
}

func Economy(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello Economy!"))
}
