package impl

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/drhelius/grpc-demo-order/internal/clients"
	"github.com/drhelius/grpc-demo-proto/order"
	"github.com/drhelius/grpc-demo-proto/product"
)

type Server struct {
	order.UnimplementedOrderServiceServer
}

func (s *Server) Create(ctx context.Context, in *order.CreateOrderReq) (*order.CreateOrderResp, error) {

	log.Printf("[Order] Create Req: %v", in.GetOrder())

	r := &order.CreateOrderResp{Id: strconv.Itoa(randomdata.Number(1000000))}

	log.Printf("[Order] Create Res: %v", r.GetId())

	return r, nil
}

func (s *Server) Read(ctx context.Context, in *order.ReadOrderReq) (*order.ReadOrderResp, error) {

	log.Printf("[Order] Read Req: %v", in.GetId())

	p1 := getProduct(in.GetId())
	p2 := getProduct(in.GetId())
	p3 := getProduct(in.GetId())

	var products = []*product.Product{p1, p2, p3}

	r := &order.ReadOrderResp{Order: &order.Order{Id: in.GetId(), Name: randomdata.SillyName(), Date: int64(randomdata.Number(1000000)), Products: products}}

	log.Printf("[Order] Read Res: %v", r.GetOrder())

	return r, nil
}

func getProduct(id string) *product.Product {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("[Order] Inoking Product service: %s", id)

	p, err := clients.ProductService.Read(ctx, &product.ReadProductReq{Id: id})

	if err != nil {
		log.Fatalf("[Order] Could not invoke Product service: %v", err)
	}

	log.Printf("[Order] Product service invocation: %v", p.GetProduct())

	return p.GetProduct()
}
