package postgre

import (
	"os"
	"restapi-altera/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgre() *gorm.DB {

	dsn := os.Getenv("POSTGRES_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect postgres")
	}

	err = db.AutoMigrate(&model.User{}, &model.Shoes{})
	if err != nil {
		panic(err.Error())
	}

	return db

}
