package model

import "errors"

var (
	errInvalidName     = errors.New("invalid user name")
	errInvalidEmail    = errors.New("invalid email address")
	errInvalidPassword = errors.New("invalid password")
)

// User is data model for user
type User struct {
	ID       int    `json:"id" example:"20"`
	Name     string `json:"name" example:"zhangsan"`
	Email    string `json:"email" example:"def@qq.com"`
	Password string `json:"password" example:"secret"`
}

// Validate check param validation
func (u User) Validate() error {
	if len(u.Name) <= 0 {
		return errInvalidName
	}
	if len(u.Password) <= 0 {
		return errInvalidPassword
	}
	if len(u.Email) <= 0 {
		return errInvalidEmail
	}

	return nil
}

// UserList is a list of user info
type UserList struct {
	PageIndex  int     `json:"pageIndex"`
	PageSize   int     `json:"pageSize"`
	TotalPages int     `json:"totalPages"`
	Data       []*User `json:"data"`
}
