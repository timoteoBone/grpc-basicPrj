package server

import (
	"context"

	protos "github.com/timoteoBone/grpc-basicPrj/fitness"
)

type Server struct{}

func (s *Server) SendRoutine(ctx *context.Context, ud *protos.UserData) (*protos.DietResponse, error) {
	return &protos.DietResponse{Basis: "vegetables", Calories: 3500}, nil
}
