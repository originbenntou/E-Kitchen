package main

import (
	"context"
	"fmt"

	pbUser "github.com/originbenntou/E-Kitchen/proto/user"
)

type UserService struct {
}

func (s *UserService) CreateUser(ctx context.Context, req *pbUser.CreateUserRequest) (*pbUser.CreateUserResponce, error) {
	fmt.Println(ctx, req)
	return nil,nil
}
