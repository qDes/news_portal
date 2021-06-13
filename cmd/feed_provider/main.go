package main

import (
	"log"
	"net"
	feed "news_portal/api/proto"
	"news_portal/internal"
	feedServer "news_portal/internal/app/feed"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger := internal.InitLogger()
	zap.ReplaceGlobals(logger)

	dbConn := "postgresql://user:pass@0.0.0.0:5432/postgres?sslmode=disable"
	port := "11000"
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	feed.RegisterFeedServiceServer(grpcServer, feedServer.NewFeedSrv(dbConn))
	grpcServer.Serve(lis)

}
