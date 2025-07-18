package tastes

import (
	"clean_code/internal/ports"

	"gorm.io/gorm"
)

var _ ports.TasteRepository = &TasteRepository{}

type TasteRepository struct {
	DB *gorm.DB
}