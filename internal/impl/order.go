package impl

import (
	"context"
	"log"
	"time"

	"github.com/drhelius/grpc-demo-order/internal/clients"
	"github.com/drhelius/grpc-demo-proto/order"
	"github.com/drhelius/grpc-demo-proto/product"
)

type Server struct {
	order.UnimplementedOrderServiceServer
}

func (s *Server) Create(ctx context.Context, in *order.CreateOrderReq) (*order.CreateOrderResp, error) {

	log.Printf("[Order] Received: %s", in.GetOrder())

	return &order.CreateOrderResp{Id: "testid"}, nil
}

func (s *Server) Read(ctx context.Context, in *order.ReadOrderReq) (*order.ReadOrderResp, error) {

	log.Printf("[Order] Received: %v", in.GetId())

	p := getProduct(in.GetId())

	var products = []*product.Product{
		{
			Id:          "001",
			Name:        "one",
			Description: "desc one",
			Price:       100,
		},
		{
			Id:          "002",
			Name:        "two",
			Description: "desc two",
			Price:       200,
		},
		p,
	}

	return &order.ReadOrderResp{Order: &order.Order{Id: "demoid", Name: "demoname", Date: 4000000, Products: products}}, nil
}

func getProduct(id string) *product.Product {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	p, err := clients.ProductService.Read(ctx, &product.ReadProductReq{Id: "000"})

	if err != nil {
		log.Fatalf("[Order] Could not invoke Product service: %v", err)
	}

	log.Printf("[Order] Product service invocation: %v", p.GetProduct())

	return p.GetProduct()
}