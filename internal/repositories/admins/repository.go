package admins

import (
	"clean_code/internal/ports"

	"gorm.io/gorm"
)

var _ ports.AdminRepository = &AdminRepository{}

type AdminRepository struct {
	DB *gorm.DB
}
