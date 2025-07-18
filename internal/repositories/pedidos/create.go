package repositories

import "clean_code/internal/domain"

func (r *PedidosRepository) Create(pedido *[]domain.Pedidos) (*[]domain.Pedidos, error) {
	err := r.DB.Create(&pedido).Error
	if err != nil {
		return nil, err
	}
	return pedido, nil
}
