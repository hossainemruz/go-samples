package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/hossainemruz/go-samples/grpc/urinary/gen/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultAddress = "localhost:1234"
)

var (
	addr = flag.String("addr", defaultAddress, "Address of the GRPC server")
	name = flag.String("name", "foo", "Name to greet")
)

func main() {
	// parse the flags
	flag.Parse()

	// establish a connection with the grpc server
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to establish a connection to the server: %s", err)
	}

	// create a greeter client
	c := pb.NewGreeterServiceClient(conn)

	// create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// send the request
	resp, err := c.Greet(ctx, &pb.GreetRequest{Name: *name})
	if err != nil {
		log.Fatalf("failed to send request: %s", err)
	}
	log.Println(resp.GetMessage())
}
