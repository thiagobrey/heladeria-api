package admins

import (
	"clean_code/internal/domain"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) error {
	var count int64
	db.Model(&domain.Admin{}).Where("email = ?", "admin@admin.com").Count(&count)
	if count == 0 {
		admin := domain.Admin{
			Email:    "admin@admin.com",
			Password: "admin123", 
			Phone:   "",
			Birthdate: "",
		}
		return db.Create(&admin).Error
	}
	return nil
}
