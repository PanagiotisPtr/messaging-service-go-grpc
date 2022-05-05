package main

import (
	"log"
	"net"
	"os"
	"panagiotisptr/messaging/protos"
	"panagiotisptr/messaging/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := log.New(os.Stdout, "messaging-service", log.Lshortfile)

	gs := grpc.NewServer()
	cs := server.NewMessagingServer(logger)

	protos.RegisterMessagingServer(gs, cs)

	reflection.Register(gs)

	list, err := net.Listen("tcp", ":9090")
	if err != nil {
		logger.Fatalf("Unable to listen. Error: %v", err)
		os.Exit(1)
	}

	gs.Serve(list)
}
