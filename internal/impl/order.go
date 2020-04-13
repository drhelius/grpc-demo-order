package impl

import (
	"context"
	"log"

	pb "github.com/drhelius/grpc-demo-proto/order"
	"github.com/drhelius/grpc-demo-proto/product"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
}

func (s *Server) Create(ctx context.Context, in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {

	log.Printf("Received: %s", in.GetOrder())

	return &pb.CreateOrderResp{Id: "testid"}, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadOrderReq) (*pb.ReadOrderResp, error) {

	log.Printf("Received: %v", in.GetId())

	var products = []*product.Product{
		&product.Product{
			Id:          "001",
			Name:        "one",
			Description: "desc one",
			Price:       100,
		},
		&product.Product{
			Id:          "002",
			Name:        "two",
			Description: "desc two",
			Price:       200,
		},
	}

	return &pb.ReadOrderResp{Order: &pb.Order{Id: "demoid", Name: "demoname", Date: 4000000, Products: products}}, nil
}
