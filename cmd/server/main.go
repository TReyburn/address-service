package main

import (
	"log"
	"net"

	"github.com/TReyburn/address-service/addresspb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	log.Println("Server starting up")

	lis, err := net.Listen("tcp", "0.0.0.0.50051")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	addresspb.RegisterAddressParseService()
}
