package main

import (
	"github.com/gorilla/mux"
	"github.com/originbenntou/E-Kitchen/front/handler"
	"github.com/originbenntou/E-Kitchen/front/middleware"
	"github.com/originbenntou/E-Kitchen/front/session"
	pbUser "github.com/originbenntou/E-Kitchen/proto/user"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	r := mux.NewRouter()

	// FIXME: ConfigMap
	userClient := pbUser.NewUserServiceClient(grpcConnect("e-kitchen-user:50051"))
	sessionStore := session.NewStoreOnMemory()

	// FrontServerは各マイクロサービスにDial
	fs := &handler.FrontServer{
		UserClient:   userClient,
		SessionStore: sessionStore,
	}

	r.Use(middleware.Logging)

	auth := middleware.NewAuthentication(userClient, sessionStore)

	r.Path("/").Methods(http.MethodGet).HandlerFunc(auth(fs.IndexHandler))
	r.Path("/signin").Methods(http.MethodGet).HandlerFunc(fs.SigninHandler)
	r.Path("/user-verify").Methods(http.MethodPost).HandlerFunc(fs.UserVerifyHandler)
	r.Path("/signup").Methods(http.MethodGet).HandlerFunc(fs.SignupHandler)
	r.Path("/user-regist").Methods(http.MethodPost).HandlerFunc(fs.UserRegistHandler)
	r.Path("/signout").Methods(http.MethodGet).HandlerFunc(auth(fs.SignoutHandler))
	r.Path("/health-check").Methods(http.MethodGet).HandlerFunc(fs.HealthCheckHandler)

	http.Handle("/", r)

	log.Println("start server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func grpcConnect(target string) *grpc.ClientConn {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn
}
