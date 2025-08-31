package admins

import (
	"clean_code/hashpassword"
	"clean_code/internal/domain"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) error {
	var count int64
	db.Model(&domain.Admin{}).Where("email = ?", os.Getenv("USER_ADMIN")).Count(&count)

	hashed, err := hashpassword.PasswordHashing(os.Getenv("USER_PASS"))
	if err != nil {
		return err
	}
	fmt.Println("Email a guardar:", os.Getenv("USER_ADMIN"))

	if count == 0 {
		admin := domain.Admin{
			Email:     os.Getenv("USER_ADMIN"),
			Password:  hashed,
			Phone:     "",
			Birthdate: "",
		}
		return db.Create(&admin).Error
	}
	return nil
}
