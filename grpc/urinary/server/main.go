package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/hossainemruz/go-samples/grpc/urinary/gen/api"
	"google.golang.org/grpc"
)

const (
	defaultPort = 1234
)

var (
	port = flag.Int("port", defaultPort, "The port to listent at")
)

type server struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *server) Greet(_ context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greeting %s", req.GetName())

	return &pb.GreetResponse{
		Message: "Hello " + req.GetName(),
	}, nil
}

func main() {
	// parse the flag
	flag.Parse()

	// create a listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to create listener: %s", err)
	}

	// register greeter server
	s := grpc.NewServer()
	pb.RegisterGreeterServiceServer(s, &server{})

	// start the server
	log.Printf("Starting server at: %s", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("fialed to start server: %s", err)
	}
	log.Println("Server stopped.")
}
