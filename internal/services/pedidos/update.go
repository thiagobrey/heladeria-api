package services

import "clean_code/internal/domain"

func (s *PedidosServices) Update(pedido *domain.Pedidos) (*domain.Pedidos, error) {
	_, err := s.RepoUser.GetUserByID(pedido.UserId)
	if err != nil {
		return nil, err
	}

	updatedPedido, err := s.RepoPedidos.Update(pedido)
	if err != nil {
		return nil, err
	}

	return updatedPedido, nil
}
