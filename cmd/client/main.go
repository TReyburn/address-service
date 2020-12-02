package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/TReyburn/address-service/addresspb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect:", err)
	}

	defer conn.Close()

	client := addresspb.NewAddressParseServiceClient(conn)

	fh := getFileHandle(os.Args[1])
	addrs := parseFile(fh)

	for _, addr := range addrs {
		callAddressParse(client, addr)
	}
	log.Println("Fin")
}

func callAddressParse(c addresspb.AddressParseServiceClient, addr string) {
	req := addresspb.APRequest{Address: addr}

	_, err := c.AddressParse(context.Background(), &req)
	if err != nil {
		log.Fatalln("Failed to call AddressParse", err)
	}

	// log.Println(resp)
}

func parseFile(fh *os.File) []string {
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	content := make([]string, 0)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content
}

func getFileHandle(fn string) *os.File {
	fh, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return fh
}
