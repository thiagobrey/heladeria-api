package admins

import (
	"clean_code/internal/domain"
	"os"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) error {
	var count int64
	db.Model(&domain.Admin{}).Where("email = ?", os.Getenv("USER_ADMIN")).Count(&count)
	if count == 0 {
		admin := domain.Admin{
			Email:     os.Getenv("USER_ADMIN"),
			Password:  os.Getenv("USER_PASS"),
			Phone:     "",
			Birthdate: "",
		}
		return db.Create(&admin).Error
	}
	return nil
}
