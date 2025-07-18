package repositories

import (
	"clean_code/internal/domain"
	"errors"

	"gorm.io/gorm"
)

func (r *PedidosRepository) GetById(id uint) (bool, error) {
	pedido := &domain.Pedidos{}
	err := r.DB.First(&pedido, "ID = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}

	if pedido.ID > 0 {
		return false, errors.New("comment already exists")
	}

	return true, nil
}
