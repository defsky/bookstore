package model

import (
	"errors"
	"fmt"
	"math"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/defsky/bookstore/basic/db"
)

// User user model
type User struct {
	gorm.Model `json:"gormModel"`
	Email      string `gorm:"unique_index;not null" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	Name       string `json:"name"`
}

// UserList ...
type UserList struct {
	PageIndex, PageSize, TotalPages int32
	Data                            []*User
}

// UserRepo ...
type UserRepo struct {
	db *gorm.DB
}

// GetUserRepo ...
func GetUserRepo() *UserRepo {
	return &UserRepo{
		db: db.GetConn(),
	}
}

// Migrate ...
func (repo *UserRepo) Migrate() {
	repo.db.AutoMigrate(&User{})
}

// Create ...
func (repo *UserRepo) Create(u *User) (*User, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(u.Password), 16)
	if err != nil {
		return nil, err
	}
	u.Password = string(h)

	if err := repo.db.Create(u).Error; err != nil {
		return nil, err
	}

	return repo.GetUserByEmail(u.Email)
}

// GetUserByEmailAndPassword ...
func (repo *UserRepo) GetUserByEmailAndPassword(email, password string) (*User, error) {
	u := &User{}

	if repo.db.Where("email=?", email).First(u).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}

	return u, nil
}

// GetUserByEmail ...
func (repo *UserRepo) GetUserByEmail(email string) (*User, error) {
	u := &User{}
	if repo.db.Where("email=?", email).First(u).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return u, nil
}

// GetUserList ...
//  n page number
//  s page size
func (repo *UserRepo) GetUserList(n, s int32) (*UserList, error) {
	userdb := repo.db.Model(&User{})
	if userdb.Error != nil {
		return nil, userdb.Error
	}
	var count int32
	if err := userdb.Count(&count).Error; err != nil {
		return nil, err
	}

	if n == 0 && s == 0 {
		n = 1
		s = count
	} else if n == 0 {
		n = 1
	} else if s == 0 {
		s = 15
	}

	pageindex := n
	pagesize := s
	totalPages := int32(math.Ceil(float64(count) / float64(pagesize)))
	if pageindex > totalPages {
		return nil, fmt.Errorf("page not exists: PageIndex=%d", pageindex)
	}
	data := []*User{}
	err := userdb.Offset((pageindex - 1) * pagesize).
		Limit(pagesize).
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	list := &UserList{
		TotalPages: totalPages,
		PageIndex:  pageindex,
		PageSize:   pagesize,
		Data:       make([]*User, 0),
	}
	for _, u := range data {
		list.Data = append(list.Data, u)
	}
	return list, nil
}

// GetUserByID ...
func (repo *UserRepo) GetUserByID(id uint64) (*User, error) {
	u := &User{}
	if repo.db.Where("id=?", id).First(u).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return u, nil
}
