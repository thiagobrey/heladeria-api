package sesions

import (
	"clean_code/internal/ports"

	"gorm.io/gorm"
)

var _ ports.SesionsRepository = &SesionsRepository{}

type SesionsRepository struct {
	DB *gorm.DB
}
