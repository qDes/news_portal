package main

import (
	"log"
	"net"
	feed "news_portal/api/proto"
	feedServer "news_portal/internal/app/feed"

	"google.golang.org/grpc"
)

func main() {
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
