package model

// User is user model for api
type User struct {
	ID    int    `json:"id" example:"20"`
	Name  string `json:"name" example:"zhangsan"`
	Email string `json:"email" example:"def@qq.com"`
}
