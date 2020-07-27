package handler

import (
	"context"
	"errors"
	"log"

	user "github.com/defsky/bookstore/proto/user"
	"github.com/defsky/bookstore/user/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	//UserHandler handle service request
	UserHandler *handler
)

type handler struct {
	// user.UnimplementedUserServer
	repo   *model.UserRepo
	tkRepo *model.TokenRepo
}

func (h *handler) Create(ctx context.Context, req *user.User) (resp *user.Response, err error) {
	var email, passwd string

	if email, passwd = req.Email, req.Password; len(email) == 0 || len(passwd) == 0 {
		resp.Success = false

		return nil, errors.New("user-service: email or password invalid")
	}
	log.Printf("New request create user: %s %s", email, passwd)

	var u *model.User
	u, err = h.repo.GetUserByEmail(req.Email)
	if u != nil {
		err = status.Errorf(codes.AlreadyExists, "user already exists")

		log.Println(err)
		return
	}

	u, err = h.repo.Create(&model.User{
		Email:    email,
		Password: passwd,
		Name:     req.Name,
	})

	resp = &user.Response{}
	if err != nil {
		log.Println(err)

		resp.Success = false

		return
	}
	log.Printf("user created: %v", u)

	resp.User = &user.User{
		Id:          uint64(u.ID),
		Email:       u.Email,
		Password:    u.Password,
		Name:        u.Name,
		CreatedTime: uint64(u.CreatedAt.Unix()),
		UpdatedTime: uint64(u.UpdatedAt.Unix()),
	}
	resp.Success = true

	return
}

func (h *handler) Get(ctx context.Context, req *user.User) (resp *user.Response, err error) {
	return
}

func (h *handler) GetAll(ctx context.Context, req *user.Request) (resp *user.Response, err error) {
	return
}

func (h *handler) Auth(ctx context.Context, req *user.User) (resp *user.Response, err error) {
	var u *model.User
	u, err = h.repo.GetUserByEmailAndPassword(req.Email, req.Password)

	resp = &user.Response{}
	if err != nil {
		resp.Success = false

		return
	}

	tk, err := h.tkRepo.Create(u)
	if err != nil {
		resp.Success = false

		return
	}

	resp.Success = true
	resp.Token = &user.Token{
		Value:   tk.Value,
		IsValid: true,
	}

	return
}

func (h *handler) ValidateToken(ctx context.Context, req *user.Token) (resp *user.Response, err error) {
	var valid bool
	valid, err = h.tkRepo.Validate(req.Value)

	resp = &user.Response{}
	resp.Success = valid
	if err != nil {
		return
	}
	return
}

func init() {
	UserHandler = &handler{
		repo:   model.GetUserRepo(),
		tkRepo: model.GetTokenRepo(),
	}
}
