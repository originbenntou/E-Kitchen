package handler

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/originbenntou/E-Kitchen/front/session"
	"github.com/originbenntou/E-Kitchen/front/support"
	"github.com/originbenntou/E-Kitchen/front/template"
	pbShop "github.com/originbenntou/E-Kitchen/proto/shop"
	pbTag "github.com/originbenntou/E-Kitchen/proto/tag"
	pbUser "github.com/originbenntou/E-Kitchen/proto/user"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type FrontServer struct {
	UserClient   pbUser.UserServiceClient
	ShopClient   pbShop.ShopServiceClient
	TagClient    pbTag.TagServiceClient
	SessionStore session.Store
}

type Content struct {
	PageName string
	User     string
	Shops    []*pbShop.Shop
	Tags     map[uint64]string
}

func (s *FrontServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var in empty.Empty

	resShop, err := s.ShopClient.FindShops(r.Context(), &in)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	tagNames := make(map[uint64]string)
	for _, shop := range resShop.Shops {
		resTag, err := s.TagClient.FindTags(r.Context(), &pbTag.FindTagsRequest{Id: shop.Id})
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/error", http.StatusFound)
			return
		}

		tagNames[shop.Id] = strings.Join(resTag.Name, " ")
	}

	template.Render(w, "index", &Content{
		PageName: "INDEX",
		User:     support.GetUserFromContext(r.Context()),
		Shops:    resShop.Shops,
		Tags:     tagNames,
	})
}

func (s *FrontServer) CreateShopHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id, _ := strconv.ParseUint(r.Form.Get("Id"), 10, 64)
	status, _ := strconv.Atoi(r.Form.Get("Status"))

	// shop登録
	resp, err := s.ShopClient.CreateShop(r.Context(), &pbShop.CreateShopRequest{
		Shop: &pbShop.Shop{
			Id:     id,
			Name:   r.Form.Get("Name"),
			Status: pbShop.Status(status),
			Url:    r.Form.Get("Url"),
			UserId: 1,
		},
	})
	if resp.Success == false || err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	// tag登録
	tagNames := strings.Split(r.Form.Get("tag"), ",")
	// エスケープ処理

	s.TagClient.RegisterTag(r.Context(), &pbTag.RegisterTagRequest{
		Id:   id,
		Name: tagNames,
	})

	http.Redirect(w, r, "/", http.StatusFound)
}

func (s *FrontServer) EditShopHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id, _ := strconv.ParseUint(r.Form.Get("Id"), 10, 64)
	status, _ := strconv.Atoi(r.Form.Get("Status"))

	resp, err := s.ShopClient.UpdateShop(r.Context(), &pbShop.UpdateShopRequest{
		Shop: &pbShop.Shop{
			Id:     id,
			Name:   r.Form.Get("Name"),
			Status: pbShop.Status(status),
			Url:    r.Form.Get("Url"),
			UserId: 1,
		},
	})
	if resp.Success == false || err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (s *FrontServer) DeleteShopHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id, _ := strconv.ParseUint(r.Form.Get("Id"), 10, 64)

	resp, err := s.ShopClient.DeleteShop(r.Context(), &pbShop.DeleteShopRequest{
		Shop: &pbShop.Shop{
			Id:     id,
			UserId: 1,
		},
	})
	if resp.Success == false || err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)

}

func (s *FrontServer) TagRegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var tagNames []string
	json.Unmarshal([]byte(r.Form.Get("data")), &tagNames)

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
