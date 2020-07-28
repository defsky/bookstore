package db

import (
	"log"

	"github.com/defsky/bookstore/basic/config"
	"github.com/jinzhu/gorm"

	// db drivers
	//_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbconn *gorm.DB

// GetConn return db connection object
func GetConn() *gorm.DB {
	if dbconn == nil {
		log.Fatalln("database connection not init ...")
	}
	return dbconn
}

func init() {
	dbCfg := config.GetConfig().DB
	db, err := gorm.Open(dbCfg.Driver, dbCfg.URL)
	if err != nil {

		panic(err)
	}
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(dbCfg.MaxIdleConnection)
	db.DB().SetMaxOpenConns(dbCfg.MaxOpenConnection)
	db.DB().SetConnMaxLifetime(dbCfg.ConnMaxLifetime)

	dbconn = db
}
