package server

import (
  "log"
  "golang.org/x/net/context"
  "api"
)
// Server represents the gRPC server
type ServerInterface interface {
	SayHello(ctx context.Context, in *api.RequestMessage) (*api.ResponseMessage, error)
}

type Server struct {
	ServerInterface
}

func (s *Server) SayHello(ctx context.Context, in *api.RequestMessage) (*api.ResponseMessage, error) {
  log.Printf("[SERVER] Received message from client - %s", in.Message)
  return &api.ResponseMessage{Message: "No.... no handshake. Corona time"}, nil
}
