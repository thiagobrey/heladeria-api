package repositories

import "clean_code/internal/domain"

func (r *PedidosRepository) Delete(id int) (*domain.Pedidos, error) {
	pedido := &domain.Pedidos{}
	err := r.DB.Delete(&pedido, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return pedido, nil
}
