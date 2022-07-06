package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunGRPCServer(addr string, registerServer func(*grpc.Server)) {
	grpcServer := grpc.NewServer()
	registerServer(grpcServer)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panicf("failed to listen at %s: %v", addr, err)
	}

	log.Fatal(grpcServer.Serve(lis))
}
