package feedServer

import (
	"context"
	feed "news_portal/api/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *feedSrv) GetFeed(ctx context.Context, req *feed.GetFeedRequest) (*feed.GetFeedResponse, error) {
	var res []*feed.News

	return &feed.GetFeedResponse{Feed: res}, nil
}

func (s *feedSrv) SubscribeUser(ctx context.Context, req *feed.SubscribeUserRequest) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *feedSrv) UnSubscribeUser(ctx context.Context, req *feed.UnSubscribeUserRequest) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}
