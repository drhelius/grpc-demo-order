package clients

import (
	"log"

	"github.com/drhelius/grpc-demo-proto/product"
	"google.golang.org/grpc"
)

var ProductService product.ProductServiceClient

func init() {
	conn, err := grpc.Dial("product:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[Order] Product client did not connect: %v", err)
	}

	ProductService = product.NewProductServiceClient(conn)
}
