package main

import (
	"context"
	pbTag "github.com/originbenntou/E-Kitchen/proto/tag"
	"github.com/originbenntou/E-Kitchen/shared/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"time"
)

type TagService struct {
	db *db.GormMutex
}

type Tag struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Status    uint64    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TagMap struct {
	Id        uint64    `json:"id"`
	ShopId    uint64    `json:"shop_id"`
	TagId     uint64    `json:"tag_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newTagGormMutex() *db.GormMutex {
	return &db.GormMutex{
		// FIXME: ConfigMap,Secretへ移行
		DBMS:   "mysql",
		USER:   "root",
		PASS:   "password",
		DBHOST: "e-kitchen-mysql:3306",
		DBNAME: "resource",
		OPTION: "charset=utf8&parseTime=True",
	}
}

func (s *TagService) FindTags(ctx context.Context, req *pbTag.FindTagsRequest) (*pbTag.FindTagsResponse, error) {
	var tagMaps []*TagMap
	if errList := s.db.SelectByWhereOneColumn(&tagMaps, "shop_id", strconv.FormatUint(req.Id, 10)).GetErrors(); len(errList) > 0 {
		for _, err := range errList {
			log.Printf("find shops failed: %s", err)
		}
		return nil, status.Error(codes.Internal, "Server Error")
	}

	var tagIds []string
	for _, tagMap := range tagMaps {
		tagIds = append(tagIds, strconv.FormatUint(tagMap.TagId, 10))
	}

	var tags []*Tag
	if errList := s.db.SelectByWhereIn(&tags, "id", tagIds).GetErrors(); len(errList) > 0 {
		for _, err := range errList {
			log.Printf("find shops failed: %s", err)
		}
		return nil, status.Error(codes.Internal, "Server Error")
	}

	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	return &pbTag.FindTagsResponse{Name: tagNames}, nil
}
