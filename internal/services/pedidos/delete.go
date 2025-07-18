package services

import "clean_code/internal/domain"

func (s *PedidosServices) Delete(id int) (*domain.Pedidos, error) {
	pedido, err := s.RepoPedidos.Delete(id)
	if err != nil {
		return nil, err
	}

	return pedido, nil
}
