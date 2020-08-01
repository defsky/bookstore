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
	log.Printf("request for create: %v", req)

	var email, passwd string
	if email, passwd = req.Email, req.Password; len(email) == 0 || len(passwd) == 0 {
		resp.Success = false

		return nil, status.Error(codes.InvalidArgument, "email or password invalid")
	}

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
		err = status.Errorf(codes.Internal, "user create failed: %s", err)
		log.Println(err)

		resp.Success = false

		return
	}
	log.Printf("response for create: %v", u)

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
	log.Printf("request for get: %v", req)

	var u *model.User
	resp = &user.Response{}

	if req.Id > 0 {
		u, err = h.repo.GetUserByID(req.Id)
	} else if len(req.Email) > 0 {
		u, err = h.repo.GetUserByEmail(req.Email)
	} else {
		err = errors.New("need user id or email")
	}
	if err != nil {
		err = status.Errorf(codes.Internal, "user get failed: %s", err)
		log.Println(err)

		resp.Success = false

		return
	}

	log.Printf("response for get: %v", u)

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

func (h *handler) GetAll(ctx context.Context, req *user.Request) (resp *user.Response, err error) {
	log.Printf("request for GetAll: %v", req)
	return
}

func (h *handler) Auth(ctx context.Context, req *user.User) (resp *user.Response, err error) {
	log.Printf("request for auth: %v", req)

	resp = &user.Response{}
	var u *model.User

	if u, err = h.repo.GetUserByEmailAndPassword(req.Email, req.Password); err == nil {
		tk, er := h.tkRepo.Create(u)
		if er == nil {
			resp.Success = true
			resp.Token = &user.Token{
				Value:   tk.Value,
				IsValid: true,
			}

			log.Printf("response for auth: %v", resp.Token)
			return
		}
		err = er
	}

	err = status.Errorf(codes.Internal, "user auth failed: %s", err)
	resp.Success = false

	return
}

func (h *handler) ValidateToken(ctx context.Context, req *user.Token) (resp *user.Response, err error) {
	log.Printf("request for validate: %v", req)
	var valid bool
	valid, err = h.tkRepo.Validate(req.Value)
	if err != nil {
		err = status.Errorf(codes.Internal, "validate failed: %s", err)
	}
	resp = &user.Response{}
	resp.Success = valid

	log.Printf("response for validate: %v", valid)

	return
}

func init() {
	UserHandler = &handler{
		repo:   model.GetUserRepo(),
		tkRepo: model.GetTokenRepo(),
	}
}
