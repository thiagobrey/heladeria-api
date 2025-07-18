package services

import "clean_code/internal/domain"

func (s *PedidosServices) List() ([]*domain.Pedidos, error) {
	pedido, err := s.RepoPedidos.List()
	if err != nil {
		return nil, err
	}

	return pedido, nil
}
