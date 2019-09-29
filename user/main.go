package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pbUser "github.com/originbenntou/E-Kitchen/proto/user"
)

const port = ":50051"

func main() {
	server := grpc.NewServer()
	pbUser.RegisterUserServiceServer(server, &UserService{})

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to create listener: %s", err)
	}

	log.Println("start server on port", port)

	if err := server.Serve(listener); err != nil {
		log.Println("failed to serve: ", err)
	}
}
