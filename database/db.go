package database

import (
	"clean_code/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConection() (*gorm.DB, error) {
	dsn := "root:301205@tcp(localhost:3306)/heladeria?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate( &domain.User{}, &domain.Taste{}, &domain.Cantidad{})

	return db, err
}
