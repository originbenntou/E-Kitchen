package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/originbenntou/E-Kitchen/front/handler"
	"github.com/originbenntou/E-Kitchen/front/middleware"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"

	pbUser "github.com/originbenntou/E-Kitchen/proto/user"
)

const port = ":8080"

func main() {
	r := mux.NewRouter()

	r.Use(middleware.Logging)

	// TODO:関数に切り出す
	target := "e-kitchen-user:50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	client := pbUser.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.CreateUser(ctx, &pbUser.CreateUserRequest{
		Email:    "hogehoge",
		Password: []byte{},
	})

	log.Println(resp)

	r.Path("/").Methods(http.MethodGet).HandlerFunc(handler.LoginHandler)
	r.Path("/signup").Methods(http.MethodGet).HandlerFunc(handler.SignupHandler)
	r.Path("/user-register").Methods(http.MethodPost).HandlerFunc(handler.UserRegisterHandler)
	r.Path("/home").Methods(http.MethodGet).HandlerFunc(handler.HomeHandler)
	r.Path("/health-check").Methods(http.MethodGet).HandlerFunc(handler.HealthCheckHandler)

	http.Handle("/", r)

	log.Println("start server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
