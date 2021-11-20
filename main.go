package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"

	"hello"

	"google.golang.org/grpc"
)

type Config struct {
	Port string `json:"port"`
}

func getConfig() Config {
	var config Config
	file, err := os.Open("config.json")

	defer file.Close()

	if err != nil {
		panic("failed to load configuration")
	}

	json.NewDecoder(file).Decode(&config)

	return config
}

func main() {
	config := getConfig()

	lis, err := net.Listen("tcp", ":"+config.Port)

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	hello.RegisterHelloServiceServer(server, &helloServiceServer{})

	log.Printf("Listening for requests ...")

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type helloServiceServer struct {
	hello.HelloServiceServer
}

func (s *helloServiceServer) GetMessage(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{
		Message: "Hello, gRPC!",
	}, nil
}
