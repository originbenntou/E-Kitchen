package handler

import (
	"fmt"
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
	Shops    []*pbShop.Shop
}

func (s *FrontServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var in empty.Empty

	res, err := s.ShopClient.FindShops(r.Context(), &in)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	template.Render(w, "index", &Content{PageName: "INDEX", Shops: res.Shops})
}

func (s *FrontServer) EditShopHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form.Get("Name"), r.Form.Get("Status"))

	resp, err := s.ShopClient.UpdateShop(r.Context(), &pbShop.UpdateShopRequest{
		Shop: &pbShop.Shop{
			Name:   r.Form.Get("Name"),
			Status: pbShop.Status(pbShop.Status_value[r.Form.Get("Status")]),
		},
	})
	if resp.Success == false || err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
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
