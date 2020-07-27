package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/defsky/bookstore/basic/db"
)

// User user model
type User struct {
	gorm.Model
	Email    string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Name     string
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
		return nil, errors.New("user not found in db")
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
		return nil, errors.New("user not found in db")
	}

	return u, nil
}
