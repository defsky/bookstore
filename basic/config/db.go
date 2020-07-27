package config

import (
	"time"
)

var defaultDBConfig = []byte(`{
	"driver":"mysql",
	"url":"root:root@(127.0.0.1:3306)/micro_book_mall?charset=utf8&parseTime=true&loc=Asia%2FShanghai",
	"maxIdleConnection":100,
	"maxOpenConnection":130,
	"connMaxLifetime":100
}`)

// DBConfig store db config data
//
// Environment Variables:
//
//   DB_DRIVER
//   DB_URL
//   DB_MAXIDLECONN
//   DB_MAXOPENCONN
//   DB_CONNMAXLIFE
type DBConfig struct {
	Driver            string        `json:"driver"`
	URL               string        `json:"url"`
	MaxIdleConnection int           `json:"maxIdleConnection"`
	MaxOpenConnection int           `json:"maxOpenConnection"`
	ConnMaxLifetime   time.Duration `json:"connMaxLifetime"`
}
