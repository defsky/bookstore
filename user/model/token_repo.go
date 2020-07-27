package model

import (
	"errors"
	"time"

	"github.com/defsky/bookstore/basic/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

var myJWTSecret = []byte("my jwt secret")

// Token user token model
type Token struct {
	gorm.Model
	Value string
}

type customClaim struct {
	jwt.StandardClaims
	uid uint
}

// TokenRepo ...
type TokenRepo struct {
	db *gorm.DB
}

// GetTokenRepo ...
func GetTokenRepo() *TokenRepo {
	return &TokenRepo{
		db: db.GetConn(),
	}
}

// Create ...
func (repo *TokenRepo) Create(u *User) (*Token, error) {
	claim := &customClaim{
		uid: u.ID,
	}

	now := time.Now()
	claim.IssuedAt = now.Unix()
	claim.Issuer = "afkplayer.com"
	// expires after 7 days
	claim.ExpiresAt = now.Add(time.Hour * 24 * 7).Unix()

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tk.Header["kid"] = u.ID

	v, err := tk.SignedString(u.Password)
	if err != nil {
		return nil, err
	}
	return &Token{
		Value: v,
	}, nil
}

// Delete ...
func (repo *TokenRepo) Delete(t *Token) bool {
	return true
}

// Validate ...
func (repo *TokenRepo) Validate(token string) (bool, error) {

	tk, err := jwt.ParseWithClaims(token,
		&customClaim{},
		func(token *jwt.Token) (interface{}, error) {
			kid := token.Header["kid"].(uint)
			u := &User{}
			if repo.db.Where("id=?", kid).First(u).RecordNotFound() {
				return nil, UserNotFound
			}

			return u.Password, nil
		})

	if err != nil {
		return false, err
	}
	if _, ok := tk.Claims.(*customClaim); ok && tk.Valid {
		return true, nil
	}
	return false, errors.New("invalid token")
}
