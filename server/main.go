package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/pengswift/planet/planet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":60001"
)

type server struct {
}

func (s *server) init() {

}

func (s *server) PrayerThrows(stream pb.PlanetService_PrayerThrowsServer) error {
	md, ok := metadata.FromContext(stream.Context())
	if !ok {
		return errors.New("")
	}

	fmt.Printf(" md[key1] :%s\n", md["key1"][0])
	fmt.Printf(" md[key2] :%s\n", md["key2"][0])

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
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ins := &server{}
	ins.init()

	pb.RegisterPlanetServiceServer(s, ins)
	s.Serve(lis)
}
