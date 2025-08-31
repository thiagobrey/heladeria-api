package repositories

import (
	"clean_code/internal/domain"
	"fmt"
)

func (r *PedidosRepository) Update(pedido *domain.Pedidos) (*domain.Pedidos, error) {
	fmt.Println("Editing comment:", pedido)
	err := r.DB.Model(&domain.Pedidos{}).
		Where("id = ? and user_id = ?", pedido.ID, pedido.UserId).
		Updates(pedido).Error
	if err != nil {
		return nil, err
	}
	return pedido, nil
}
