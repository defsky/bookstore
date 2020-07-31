package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/defsky/bookstore/micro"
	userPb "github.com/defsky/bookstore/proto/user"
)

var isCreate, isGet, isAuth, isValidate bool
var tokenStr string

func main() {
	log.Println("user-cli started")

	if len(os.Args) < 2 {
		log.Fatalf("need flags,%s -h for help", os.Args[0])
	}

	flag.BoolVar(&isCreate, "create", false, "create user info or not")
	flag.BoolVar(&isGet, "get", false, "get user info or not")
	flag.BoolVar(&isAuth, "auth", false, "auth user info or not")
	flag.BoolVar(&isValidate, "validate", false, "validate user info or not")
	flag.StringVar(&tokenStr, "token", "", "token string")
	flag.Parse()

	c := micro.NewClient(
		micro.Name("com.afkplayer.service.user"),
	)
	client := userPb.NewUserClient(c.Conn())

	user := &userPb.User{
		Name:     "def",
		Email:    "defsky@qq.com",
		Password: "123",
	}

	if isCreate {
		resp, err := client.Create(context.Background(), user)
		if err != nil {
			log.Fatalln(err)
		}
		if resp.Success {
			log.Printf("user created: %v", resp.User)
		}
	}
	if isGet {
		resp, err := client.Get(context.Background(), user)
		if err != nil {
			log.Fatalln(err)
		}
		if resp.Success {
			log.Printf("user info: %v", resp.User)
		}
		return
	}
	if isAuth {
		resp, err := client.Auth(context.Background(), user)
		if err != nil {
			log.Fatalln(err)
		}
		if resp.Success {
			log.Printf("user token: %v", resp.Token)
		}
	}
	if isValidate {
		if len(tokenStr) <= 0 {
			log.Fatalln("token should not be empty string")
		}
		tk := &userPb.Token{
			Value: tokenStr,
		}
		resp, err := client.ValidateToken(context.Background(), tk)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("validate response: %v", resp)
	}
}
