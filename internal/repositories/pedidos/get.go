package repositories

import "clean_code/internal/domain"

func (r *PedidosRepository) GetByUserId(userId int) ([]*domain.Pedidos, error) {
	var pedido []*domain.Pedidos
	err := r.DB.Where("user_id = ?", userId).Find(&pedido).Error
	if err != nil {
		return nil, err
	}
	return pedido, nil
}
