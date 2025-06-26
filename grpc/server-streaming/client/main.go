package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/hossainemruz/go-samples/grpc/server-streaming/gen/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr   = flag.String("addr", "localhost:1234", "Address of the server")
	symbol = flag.String("symbol", "TSLA", "Symbol of the stock")
)

func streamPrice(ctx context.Context, client pb.StockServiceClient, symbol string) error {
	stream, err := client.StreamPrice(ctx, &pb.StreamPriceRequest{StockSymbol: symbol})
	if err != nil {
		log.Fatalf("failed to create stream: %s", err)
	}

	log.Printf("steaming price for: %s", symbol)
	for {
		update, err := stream.Recv()
		if err == io.EOF {
			log.Println("server closed the stream")
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Time: %s, Price: %f", update.GetTimestamp().AsTime(), update.GetPrice())
	}
}

func main() {
	flag.Parse()

	// establish a connection to the server
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create connection: %s", err)
	}
	defer conn.Close()

	// create stock service client
	client := pb.NewStockServiceClient(conn)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}

	// create a context so we can cancel the stream
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a signal handler so that we can cancel on keyboard interrupt
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// start a go routine that will cancel the context when termination signal is received
	go func() {
		// wait for termination sigal to come on the channel
		<-sigChan
		log.Println("cancelling context")
		cancel()
	}()

	// stream the price
	// this part is blocking
	err = streamPrice(ctx, client, *symbol)
	if err != nil {
		log.Fatalf("streaming ended with error: %s", err)
	}
}
