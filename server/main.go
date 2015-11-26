package main

import (
	"log"
	"net"

	pb "github.com/pengswift/planet/planet"
	"google.golang.org/grpc"
)

const (
	port = ":70001"
)

type server struct {
}

func (s *server) init() {

}

func (s *server) Stream(PlanetService_StreamServer) error

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ins := &server{}
	ins.init()

	pb.RegisterPlanetServiceServer(s, ins)
	s.Serve(lis)
}
