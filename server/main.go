package main

import (
	"errors"
	"io"
	"log"
	"net"

	pb "github.com/pengswift/planet/planet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":70001"
)

type server struct {
}

func (s *server) init() {

}

func (s *server) PrayerThrows(stream pb.PlanetService_PrayerThrowServer) error {
	md, ok := metadata.FromContext(stream.Context())
	if !ok {
		return errors.New("")
	}

	println(" md [key1] :%s", md["key1"])
	println(" md [key2] :%s", md["key2"])

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		_ = in

	}

	return nil
}

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
