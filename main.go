package main

import (
	"github.com/hashicorp/go-hclog"
	protos "github.com/timoteoBone/grpc-basicPrj/fitness"
	"github.com/timoteoBone/grpc-basicPrj/server"
	"google.golang.org/grpc"
)

func main() {

	log := hclog.Default()
	sv := server.NewServer(log)

	grpcSv := &grpc.Server{}

	protos.RegisterNutricionistServer(grpcSv, sv)

}
