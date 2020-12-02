package main

import (
	"context"
	"log"
	"net"

	"github.com/TReyburn/address-service/addresspb"
	"github.com/TReyburn/address-service/parser"
	"google.golang.org/grpc"
)

type server struct{}

func (s server) AddressParse(_ context.Context, req *addresspb.APRequest) (*addresspb.APResponse, error) {
	a := req.GetAddress()
	c := make(chan *addresspb.APResponse)
	go parser.ParseAddress(a, c)
	return <-c, nil
}

func main() {
	log.Println("Server starting up")
	parser.InitParser()

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	addresspb.RegisterAddressParseServiceServer(s, server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
