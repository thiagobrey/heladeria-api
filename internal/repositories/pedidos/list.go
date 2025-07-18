package repositories

import "clean_code/internal/domain"

func (r *PedidosRepository) List() ([]*domain.Pedidos, error) {
	var pedido []*domain.Pedidos
	err := r.DB.Find(&pedido).Error
	if err != nil {
		return nil, err
	}
	return pedido, nil
}
