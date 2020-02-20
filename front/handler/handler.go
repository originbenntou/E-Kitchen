package handler

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/originbenntou/E-Kitchen/front/session"
	"github.com/originbenntou/E-Kitchen/front/template"
	pbShop "github.com/originbenntou/E-Kitchen/proto/shop"
	pbUser "github.com/originbenntou/E-Kitchen/proto/user"
	"io"
	"log"
	"net/http"
)

type FrontServer struct {
	UserClient   pbUser.UserServiceClient
	ShopClient   pbShop.ShopServiceClient
	SessionStore session.Store
}

type Content struct {
	PageName string
}

func (s *FrontServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var in empty.Empty

	log.Println(s.ShopClient.FindShops(r.Context(), &in))
	template.Render(w, "index", &Content{PageName: "INDEX"})
}

func (s *FrontServer) SigninHandler(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "signin", &Content{PageName: "SIGN_IN"})
}

func (s *FrontServer) UserVerifyHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := s.UserClient.VerifyUser(r.Context(), &pbUser.VerifyUserRequest{
		Email:    r.Form.Get("email"),
		Password: []byte(r.Form.Get("password")),
	})
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	sessionID := session.ID()
	s.SessionStore.Set(sessionID, resp.User.Id)
	session.SetSessionIDToResponse(w, sessionID)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (s *FrontServer) SignupHandler(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "signup", &Content{PageName: "SIGN_UP"})
}

func (s *FrontServer) UserRegistHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := s.UserClient.CreateUser(r.Context(), &pbUser.CreateUserRequest{
		Email:    r.Form.Get("email"),
		Password: []byte(r.Form.Get("password")),
	})
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	sessionID := session.ID()
	s.SessionStore.Set(sessionID, resp.GetUser().GetId())
	session.SetSessionIDToResponse(w, sessionID)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (s *FrontServer) SignoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := session.GetSessionIDFromRequest(r)
	s.SessionStore.Delete(sessionID)
	session.DeleteSessionIDFromResponse(w)
	http.Redirect(w, r, "/signin", http.StatusFound)
}

// ヘルスチェック
func (s *FrontServer) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, `{"alive": true}`)
}

func (s *FrontServer) ErrorHandler(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "error", &Content{PageName: "error"})
}
