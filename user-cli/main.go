package main

import (
	"context"
	"log"

	"github.com/defsky/bookstore/micro"
	userPb "github.com/defsky/bookstore/proto/user"
)

func main() {
	log.Println("user-cli started")

	c := micro.NewClient(
		micro.Name("com.afkplayer.service.user"),
	)
	client := userPb.NewUserClient(c.Conn())

	user := &userPb.User{
		Name:     "def",
		Email:    "defsky@qq.com",
		Password: "123",
	}
	resp, err := client.Create(context.Background(), user)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Success {
		log.Printf("user created: %v", resp.User)
	}
}
