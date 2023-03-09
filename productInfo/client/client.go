package main

import (
	"github.com/linqcod/grpc-up-and-running/productInfo/client/ecommerce"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error while connecting to the serverL %v", err)
	}
	defer conn.Close()

	c := ecommerce.NewProductInfoClient(conn)

	name := "Apple iPhone 11"
	description := `Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode.`
	price := float32(1000.0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddProduct(ctx, &ecommerce.Product{
		Name:        name,
		Description: description,
		Price:       price,
	})
	if err != nil {
		log.Fatalf("error while adding new product: %v", err)
	}
	log.Printf("Product ID: %s added successfully!", r.Value)

	product, err := c.GetProduct(ctx, r)
	if err != nil {
		log.Fatalf("error while getting product by id: %v", err)
	}
	log.Printf("Product: %s", product.String())
}
