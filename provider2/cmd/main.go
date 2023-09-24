package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"provider2/internal/config"
	"provider2/internal/providergrpc"
	"provider2/internal/server"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", config.EnvConfigs.LocalServerPort)
	if err != nil {
		log.Fatalf("Error when starting the server: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()
	providergrpc.RegisterProviderServer(s, &server.Server{})

	log.Println("gRPC server running on :3002 ...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
