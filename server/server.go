package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/timoteoBone/grpc-basicPrj/fitness"
)

type Server struct {
	log hclog.Logger
}

func NewServer(log hclog.Logger) *Server {
	return &Server{log}
}

func (s *Server) SendRoutine(con context.Context, ud *protos.UserData) (*protos.DietResponse, error) {
	s.log.Info("sd")
	return &protos.DietResponse{Basis: "vegetables", Calories: 3500}, nil
}
