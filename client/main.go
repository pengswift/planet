package main

import (
	"log"

	pb "github.com/pengswift/planet/planet"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:70001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPlanetServiceClient(conn)

	ctx = metadata.NewContext(context.Background(), metadata.Pairs("key1", "val1", "key2", "val2"))

	r, err := c.PrayerThrows(ctx, &pb.PlanetRequest{Msg: []byte("hello")})
	if err != nil {
		log.Fatalf("could not prayerthrows: %v", err)
	}
	log.Printf("receive msg: %s", string(r.Msg))
}
