package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	pbUser "github.com/originbenntou/E-Kitchen/proto/user"
	"github.com/originbenntou/E-Kitchen/shared/db"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type UserService struct {
}

type Users struct {
	Id        int       `json:"id,string"`
	Email     string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created,string"`
	UpdatedAt time.Time `json:"updated,string"`
}

func (s *UserService) VerifyUser(ctx context.Context, req *pbUser.VerifyUserRequest) (*pbUser.VerifyUserResponse, error) {
	user, errList := getUser(req.Email)
	if len(errList) > 0 {
		for _, err := range errList {
			log.Printf("verify user failed: %s", err)
		}
		return nil, status.Error(codes.Internal, "Server Error")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), req.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	existUser := &pbUser.User{
		Id:           uint64(user.Id),
		Email:        user.Email,
		PasswordHash: []byte(user.Password),
	}

	return &pbUser.VerifyUserResponse{User: existUser}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pbUser.CreateUserRequest) (*pbUser.CreateUserResponse, error) {
	if req.Email == "" || len(req.Password) < 0 {
		return nil, status.Error(codes.InvalidArgument, "empty email or password")
	}

	passHash, err := bcrypt.GenerateFromPassword(req.Password, bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	newUser := &pbUser.User{
		Email:        req.Email,
		PasswordHash: passHash,
		CreatedAt:    ptypes.TimestampNow(),
	}

	exist, errList := existUser(newUser.Email)
	if len(errList) > 0 {
		for _, err := range errList {
			log.Printf("create user failed: %s", err)
		}
		return nil, status.Error(codes.Internal, "Server Error")
	}
	if exist {
		return nil, status.Error(codes.AlreadyExists, "This User Already Exists")
	}

	if errList := createUser(newUser); len(errList) > 0 {
		for _, err := range errList {
			log.Printf("create user failed: %s", err)
		}
		return nil, status.Error(codes.Internal, "Server Error")
	}

	// sessionIDセット
	newUser.Id = getId(newUser.Email)
	if newUser.Id == 0 {
		return nil, status.Error(codes.Internal, "Server Error")
	}

	return &pbUser.CreateUserResponse{User: newUser}, nil
}

func (s *UserService) FindUser(ctx context.Context, req *pbUser.FindUserRequest) (*pbUser.FindUserResponse, error) {
	userEmail, err := findUser(req.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "User is not Found")
	}

	return &pbUser.FindUserResponse{Email: userEmail}, nil
}

func createUser(user *pbUser.User) []error {
	now, _ := ptypes.Timestamp(user.CreatedAt)
	jst, _ := time.LoadLocation("Asia/Tokyo")

	newUser := new(Users)
	newUser.Email = user.Email
	newUser.Password = string(user.PasswordHash)
	newUser.CreatedAt = now.In(jst)
	newUser.UpdatedAt = now.In(jst)

	result := db.Insert(newUser)

	return result.GetErrors()
}

func existUser(email string) (bool, []error) {
	var u Users
	r, c := db.Count(&u, "email", email)

	if err := r.GetErrors(); len(err) > 0 {
		return false, err
	}

	return c > 0, nil
}

func getId(email string) uint64 {
	var u Users
	if err := db.Select(&u, "email", email).GetErrors(); len(err) > 0 {
		return 0
	}

	return uint64(u.Id)
}

func getUser(email string) (Users, []error) {
	var u Users
	if err := db.Select(&u, "email", email).GetErrors(); len(err) > 0 {
		return Users{}, err
	}

	return u, nil
}

func findUser(id uint64) (string, []error) {
	var u Users
	if err := db.Select(&u, "id", string(id)).GetErrors(); len(err) > 0 {
		return "", err
	}

	return u.Email, nil
}
