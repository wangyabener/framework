package mysql

import (
	"fmt"
	"framework/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migration() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(models.User{})
}
