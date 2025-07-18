package services

import "clean_code/internal/domain"

func (s *PedidosServices) GetByUserId(userId int) ([]*domain.Pedidos, error) {

	_, err := s.RepoUser.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	pedido, err := s.RepoPedidos.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	return pedido, nil
}
