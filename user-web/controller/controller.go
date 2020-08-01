package controller

import (
	"context"

	"github.com/defsky/bookstore/micro"
	"github.com/defsky/bookstore/user-web/model"

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
