package config

import (
	"os"
	"restapi-altera/config/cache"
	"restapi-altera/config/postgre"
	"strconv"

	"gorm.io/gorm"
)

type Config interface {
	Database() *gorm.DB
	Redis() cache.Redis
	Port() int
}

type config struct {
}

func NewConfig() Config {
	return &config{}

}

func (c *config) Database() *gorm.DB {
	return postgre.InitPostgre()
}

func (c *config) Redis() cache.Redis {
	return cache.InitRedis()
}

func (c *config) Port() int {
	port, _ := strconv.Atoi(os.Getenv("PORT_APP"))
	return port

}
