package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"time"

	pb "github.com/hossainemruz/go-samples/grpc/server-streaming/gen/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	port = flag.Int("port", 1234, "Port where server will serve")
)

type server struct {
	pb.UnimplementedStockServiceServer
}

func (s *server) StreamPrice(req *pb.StreamPriceRequest, stream grpc.ServerStreamingServer[pb.StreamPriceResponse]) error {
	ctx := stream.Context()
	stockSymbol := req.GetStockSymbol()
	for {
		select {
		// if client cancelled the context, close streaming
		case <-ctx.Done():
			log.Printf("User cancelled the context for symbol: %s", stockSymbol)
			return ctx.Err()
		// send price update
		default:
			log.Printf("Sending price update for symbol: %s", stockSymbol)
			stream.Send(&pb.StreamPriceResponse{
				StockSymbol: stockSymbol,
				Price:       generateRandomPrice(),
				Timestamp:   timestamppb.New(time.Now()),
			})
			// sleep for 1 second to simulate some time passed
			time.Sleep(1 * time.Second)
		}

	}
}

func generateRandomPrice() float32 {
	min := float32(10.0)
	max := float32(1000.0)
	return min + rand.Float32()*(max-min)
}

func main() {
	flag.Parse()

	// create a listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to create listener: %s", err)
	}

	// create grpc server
	s := grpc.NewServer()

	// register the stock price service to the server
	pb.RegisterStockServiceServer(s, &server{})

	// start the server
	log.Printf("Starting server at: %s", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to start sever: %s", err)
	}

	log.Println("Server stopped.")
}
