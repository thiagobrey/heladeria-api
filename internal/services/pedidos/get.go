package services

import (
	"clean_code/internal/domain"
	"strconv"
	"strings"
)

func (s *PedidosServices) GetByUserId(userId int) ([]*domain.Pedidos, error) {
	_, err := s.RepoUser.GetById(userId)
	if err != nil {
		return nil, err
	}

	pedidos, err := s.RepoPedidos.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	for _, pedido := range pedidos {
		tasteIds := strings.Split(pedido.Tastes, ",")
		var tasteIdsInt []int
		for _, id := range tasteIds {
			id = strings.TrimSpace(id)
			if id == "" {
				continue
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return nil, err
			}
			tasteIdsInt = append(tasteIdsInt, idInt)
		}
		tasteNames, err := s.RepoTastes.GetByIds(tasteIdsInt)
		if err != nil {
			return nil, err
		}
		pedido.Tastes = strings.Join(tasteNames, ",")
	}

	return pedidos, nil
}
