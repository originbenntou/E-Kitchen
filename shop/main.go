package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	pbShop "github.com/originbenntou/E-Kitchen/proto/shop"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

func main() {
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)
	pbShop.RegisterShopServiceServer(server, &ShopService{
		db: newShopGormMutex(),
	})

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to create listener: %s", err)
	}

	log.Println("start server on port", port)

	if err := server.Serve(listener); err != nil {
		log.Println("failed to serve: ", err)
	}
}
