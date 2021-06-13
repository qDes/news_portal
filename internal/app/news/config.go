package news

import (
	feed "news_portal/api/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type service struct {
	FeedClient feed.FeedServiceClient
}


func GetSvc() *service {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())

	cn, err := grpc.Dial("0.0.0.0:11000", opts...)
	if err != nil {
		panic(err)
	}

	client := feed.NewFeedServiceClient(cn)
	if err != nil {
		zap.L().Error("grpc client connection error", zap.Error(err))
	}
	return &service{FeedClient: client}
}