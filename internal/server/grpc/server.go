package grpc

import (
	"log"
	"net"
	"sync"

	"github.com/drhelius/grpc-demo-order/internal/impl"
	"github.com/drhelius/grpc-demo-proto/order"
	"google.golang.org/grpc"
)

func Serve(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("[Order] GRPC failed to listen: %v", err)
	}

	s := grpc.NewServer()

	order.RegisterOrderServiceServer(s, &impl.Server{})

	log.Printf("[Order] Serving GRPC on localhost:%s ...", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("[Order] GRPC failed to serve: %v", err)
	}
}
