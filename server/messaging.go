package server

import (
	"context"
	"log"
	"panagiotisptr/messaging/protos"
)

type MessagingServer struct {
	logger *log.Logger
	protos.UnimplementedMessagingServer
}

func NewMessagingServer(l *log.Logger) *MessagingServer {
	return &MessagingServer{
		logger: l,
	}
}

func (m *MessagingServer) SendMessage(
	ctx context.Context,
	request *protos.SendMessageRequest,
) (*protos.SendMessageResponse, error) {
	return &protos.SendMessageResponse{
		Sent: true,
	}, nil
}
