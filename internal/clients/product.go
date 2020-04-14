package clients

import (
	"log"

	"github.com/drhelius/grpc-demo-proto/product"
	"google.golang.org/grpc"
)

var ProductService product.ProductServiceClient

func init() {
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	ProductService = product.NewProductServiceClient(conn)
}
