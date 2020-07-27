package config

import (
	"log"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/env"
	"github.com/micro/go-micro/v2/config/source/memory"
)

var (
	appConfig *AppConfig
)

//AppConfig ...
type AppConfig struct {
	DB *DBConfig
}

//GetConfig ...
func GetConfig() *AppConfig {
	if appConfig == nil {
		log.Fatalln("app configuration not init ...")
	}
	return appConfig
}

func init() {
	appConfig = &AppConfig{
		DB: &DBConfig{},
	}

	defaultDbCfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	defaultDbCfg.Load(memory.NewSource(memory.WithJSON(defaultDBConfig)))
	config.Scan(appConfig.DB)

	envDbCfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	envDbCfg.Load(env.NewSource(env.WithStrippedPrefix("DB")))

	if s := envDbCfg.Get("driver").String(""); s != "" {
		appConfig.DB.Driver = s
	}
	if s := envDbCfg.Get("url").String(""); s != "" {
		appConfig.DB.URL = s
	}
	if i := envDbCfg.Get("maxidleconn").Int(0); i != 0 {
		appConfig.DB.MaxIdleConnection = i
	}
	if i := envDbCfg.Get("maxopenconn").Int(0); i != 0 {
		appConfig.DB.MaxOpenConnection = i
	}
	if i := envDbCfg.Get("connmaxlife").Duration(0); i != 0 {
		appConfig.DB.ConnMaxLifetime = i
	}
}
