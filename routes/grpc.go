package routes

import "google.golang.org/grpc"

type grpcRoute struct {
	server *grpc.Server
}

func NewGrpcRoute(server *grpc.Server) *grpcRoute {
	return &grpcRoute{
		server: server,
	}
}
