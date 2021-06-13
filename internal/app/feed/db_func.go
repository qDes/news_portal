package feedServer

import (
	"context"
	feed "news_portal/api/proto"
	"news_portal/internal"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *feedSrv) GetUserFeed(ctx context.Context, req *feed.GetUserFeedRequest) (*feed.FeedResponse, error) {
	var res []*feed.News

	return &feed.FeedResponse{Feed: res}, nil
}


func (s *feedSrv) GetFeed(ctx context.Context, req *feed.GetFeedRequest) (*feed.FeedResponse, error) {
	var res []*feed.News
	news := internal.GetFeedByTopic(s.db, req.Topic)
	for _, item := range news {
		res = append(res, &feed.News{
			Id:    0,
			Title: item.Title,
			Text:  item.Text,
			Date:  item.Dttm,
		})
	}
	return &feed.FeedResponse{Feed: res}, nil
}

func (s *feedSrv) SubscribeUser(ctx context.Context, req *feed.SubscribeUserRequest) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *feedSrv) UnSubscribeUser(ctx context.Context, req *feed.UnSubscribeUserRequest) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}
