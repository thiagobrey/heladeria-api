package repositories

import (
	"clean_code/internal/ports"

	"gorm.io/gorm"
)

var _ ports.PedidosRepository = &PedidosRepository{}

type PedidosRepository struct {
	DB *gorm.DB
}
