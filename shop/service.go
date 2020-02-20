package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pbShop "github.com/originbenntou/E-Kitchen/proto/shop"
	"github.com/originbenntou/E-Kitchen/shared/db"
	"time"
)

type ShopService struct {
	db *db.GormMutex
}

type Shop struct {
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	Status     pbShop.Status `json:"status"`
	CategoryId uint64        `json:"category_id"`
	UserId     uint64        `json:"user_id"`
	CreatedAt  time.Time     `json:"created"`
	UpdatedAt  time.Time     `json:"updated"`
}

func newShopGormMutex() *db.GormMutex {
	return &db.GormMutex{
		// FIXME: ConfigMap,Secretへ移行
		DBMS:   "mysql",
		USER:   "root",
		PASS:   "password",
		DBHOST: "e-kitchen-mysql:3306",
		DBNAME: "e_kitchen",
		OPTION: "charset=utf8&parseTime=True",
	}
}

func (s *ShopService) FindShops(ctx context.Context, in *empty.Empty) (*pbShop.FindShopsResponse, error) {
	var ss []*pbShop.Shop
	if err := s.db.SelectAll(&ss).GetErrors(); len(err) > 0 {
		return &pbShop.FindShopsResponse{}, nil
	}

	fmt.Println("ここ", ss)

	return &pbShop.FindShopsResponse{Shops: ss}, nil
}
