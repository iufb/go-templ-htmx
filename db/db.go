package db

import (
	"fmt"
	"log"

	"github.com/iufb/go-templ-htmx/config"
	"github.com/iufb/go-templ-htmx/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBHost, config.Envs.DBPort, config.Envs.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to db %v", err)
	}
	db.AutoMigrate(&types.User{})
	return db
}
