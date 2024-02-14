package main

import (
	"chatapp-grpc-redis/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type chatAppServer struct {
	proto.ChatAppServer
}

func (s *chatAppServer) Add(c context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	result := req.Number + req.AnotherNumber
	response := &proto.AddResponse{
		Sum: int64(result),
	}
	log.Print(response.Sum)
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterChatAppServer(s, &chatAppServer{})
	log.Printf("Server registered successfull!!! \n")
	err2  := s.Serve(lis)
	if err2 != nil {
		log.Fatalf("err while serve %v", err)
	}
}
