package main

import (
	"context"
	"log"
	"net"

	"github.com/ahsanulks/testingrpc/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listen, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()

	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(listen); err != nil {
		panic(err)
	}
}

func (s *server) Add(context context.Context, request *proto.Request) (*proto.Response, error) {
	log.Println("have a request with params", request.X, request.Y)
	return &proto.Response{Result: request.X + request.Y}, nil
}

func (s *server) Multiply(context context.Context, request *proto.Request) (*proto.Response, error) {
	log.Println("have a request with params", request.X, request.Y)
	return &proto.Response{Result: request.X * request.Y}, nil
}
