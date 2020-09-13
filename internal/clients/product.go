package clients

import (
	"log"

	"github.com/drhelius/grpc-demo-proto/product"
	"google.golang.org/grpc"
)

var ProductService product.ProductServiceClient

func init() {
	log.Printf("[Order] Dialing to 'product:5000' ...")

	conn, err := grpc.Dial("product:5000", grpc.WithInsecure(), grpc.WithBlock(), grpc.FailOnNonTempDialError(true))
	if err != nil {
		log.Fatalf("[Order] Error dialing to Product service: %v", err)
	}

	ProductService = product.NewProductServiceClient(conn)
}
