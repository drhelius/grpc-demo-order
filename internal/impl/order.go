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

	p1 := getProduct(strconv.Itoa(randomdata.Number(1000000)))
	p2 := getProduct(strconv.Itoa(randomdata.Number(1000000)))
	p3 := getProduct(strconv.Itoa(randomdata.Number(1000000)))
	p4 := getProduct(strconv.Itoa(randomdata.Number(1000000)))
	p5 := getProduct(strconv.Itoa(randomdata.Number(1000000)))

	publicIP := clients.GetPublicIP()

	var products = []*product.Product{p1, p2, p3, p4, p5}

	r := &order.ReadOrderResp{Order: &order.Order{Id: in.GetId(), Name: randomdata.SillyName(), Date: int64(randomdata.Number(1000000)), Products: products, Ip: publicIP}}

	log.Printf("[Order] Read Res: %v", r.GetOrder())

	return r, nil
}

func getProduct(id string) *product.Product {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("[Order] Invoking Product service: %s", id)

	p, err := clients.ProductService.Read(ctx, &product.ReadProductReq{Id: id})

	if err != nil {
		log.Printf("[Order] ERROR - Could not invoke Product service: %v", err)
		return &product.Product{}
	}

	log.Printf("[Order] Product service invocation: %v", p.GetProduct())
	return p.GetProduct()
}
