package model

import "github.com/defsky/bookstore/basic/db"

func init() {
	conn := db.GetConn()

	conn.AutoMigrate(&User{})
}
