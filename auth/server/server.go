package server

import (
	"Ripped/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

type authServer struct {}

func main() {
	grpcServerListener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	grpcServer := grpc.NewServer()
	auth.RegisterAuthenticationServiceServer(grpcServer, &authServer{})

	if err := grpcServer.Serve(grpcServerListener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}