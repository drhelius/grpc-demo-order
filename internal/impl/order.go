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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	order.UnimplementedOrderServiceServer
}

func (s *Server) Create(ctx context.Context, in *order.CreateOrderReq) (*order.CreateOrderResp, error) {

	log.Printf("[Order] Create Req: %v", in.GetOrder())

	r := &order.CreateOrderResp{Id: strconv.Itoa(randomdata.Number(1000000))}

	err := failedContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("[Order] Create Res: %v", r.GetId())

	return r, nil
}

func (s *Server) Read(ctx context.Context, in *order.ReadOrderReq) (*order.ReadOrderResp, error) {

	log.Printf("[Order] Read Req: %v", in.GetId())

	err := failedContext(ctx)
	if err != nil {
		return nil, err
	}

	p1 := getProduct(ctx, strconv.Itoa(randomdata.Number(1000000)))

	err = failedContext(ctx)
	if err != nil {
		return nil, err
	}

	p2 := getProduct(ctx, strconv.Itoa(randomdata.Number(1000000)))

	err = failedContext(ctx)
	if err != nil {
		return nil, err
	}

	p3 := getProduct(ctx, strconv.Itoa(randomdata.Number(1000000)))

	err = failedContext(ctx)
	if err != nil {
		return nil, err
	}

	p4 := getProduct(ctx, strconv.Itoa(randomdata.Number(1000000)))

	err = failedContext(ctx)
	if err != nil {
		return nil, err
	}

	p5 := getProduct(ctx, strconv.Itoa(randomdata.Number(1000000)))

	err = failedContext(ctx)
	if err != nil {
		return nil, err
	}

	publicIP := clients.GetPublicIP()

	err = failedContext(ctx)
	if err != nil {
		return nil, err
	}

	var products = []*product.Product{p1, p2, p3, p4, p5}

	r := &order.ReadOrderResp{Order: &order.Order{Id: in.GetId(), Name: randomdata.SillyName(), Date: int64(randomdata.Number(1000000)), Products: products, Ip: publicIP}}

	log.Printf("[Order] Read Res: %v", r.GetOrder())

	return r, nil
}

func getProduct(ctx context.Context, id string) *product.Product {

	headersIn, _ := metadata.FromIncomingContext(ctx)

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctxTimeout, headersIn)

	log.Printf("[Order] Invoking Product service: %s", id)

	p, err := clients.ProductService.Read(ctx, &product.ReadProductReq{Id: id})

	if err != nil {
		log.Printf("[Order] ERROR - Could not invoke Product service: %v", err)
		return &product.Product{}
	}

	log.Printf("[Order] Product service invocation: %v", p.GetProduct())
	return p.GetProduct()
}

func failedContext(ctx context.Context) error {
	if ctx.Err() == context.Canceled {
		log.Printf("[Order] context canceled, stoping server side operation")
		return status.Error(codes.Canceled, "context canceled, stoping server side operation")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("[Order] dealine has exceeded, stoping server side operation")
		return status.Error(codes.DeadlineExceeded, "dealine has exceeded, stoping server side operation")
	}

	return nil
}
