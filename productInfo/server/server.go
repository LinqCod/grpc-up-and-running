package main

import (
	"github.com/gofrs/uuid"
	"github.com/linqcod/grpc-up-and-running/productInfo/server/ecommerce"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	ecommerce.UnimplementedProductInfoServer
	productMap map[string]*ecommerce.Product
}

func (s *server) AddProduct(ctx context.Context, in *ecommerce.Product) (*ecommerce.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while generating Product ID", err)
	}

	in.Id = out.String()

	if s.productMap == nil {
		s.productMap = make(map[string]*ecommerce.Product)
	}
	s.productMap[in.Id] = in

	return &ecommerce.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetProduct(ctx context.Context, in *ecommerce.ProductID) (*ecommerce.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "product with corresponding id not found, id: ", in.Value)
}
