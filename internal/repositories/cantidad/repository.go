package cantidad

import (
	"clean_code/internal/ports"

	"gorm.io/gorm"
)

var _ ports.CantidadRepository = &CantidadRepository{}

type CantidadRepository struct {
	DB *gorm.DB
}
