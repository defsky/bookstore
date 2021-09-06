package controller

import (
	"context"

	"github.com/defsky/bookstore/micro"
	"github.com/defsky/bookstore/user-api/model"

	userPb "github.com/defsky/bookstore/proto/user"
)

// Controller route controller
type Controller struct {
	userClient userPb.UserClient
}

// NewController return a new controller object
func NewController() *Controller {
	mc := micro.NewClient(
		micro.Name("com.afkplayer.service.user"),
	)
	return &Controller{
		userClient: userPb.NewUserClient(mc.Conn()),
	}
}

func (c *Controller) getUserByID(id int) (*model.User, error) {
	resp, err := c.userClient.Get(context.Background(), &userPb.User{
		Id: uint64(id),
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    int(resp.User.Id),
		Name:  resp.User.Name,
		Email: resp.User.Email,
	}, nil
}

func (c *Controller) getUserList(n, s int) (*model.UserList, error) {
	resp, err := c.userClient.GetList(context.Background(), &userPb.UserList{
		PageIndex: int32(n),
		PageSize:  int32(s),
	})
	if err != nil {
		return nil, err
	}
	list := &model.UserList{
		PageIndex:  int(resp.Users.PageIndex),
		PageSize:   int(resp.Users.PageSize),
		TotalPages: int(resp.Users.TotalPages),
		Data:       make([]*model.User, 0),
	}
	for _, u := range resp.Users.Data {
		list.Data = append(list.Data, &model.User{
			ID:       int(u.Id),
			Name:     u.Name,
			Email:    u.Email,
			Password: u.Password,
		})
	}
	return list, nil
}

func (c *Controller) createUser(u model.User) (*model.User, error) {
	resp, err := c.userClient.Create(context.Background(), &userPb.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       int(resp.User.Id),
		Email:    resp.User.Email,
		Name:     resp.User.Name,
		Password: resp.User.Password,
	}, nil
}
