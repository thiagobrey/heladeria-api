package services

import (
	"clean_code/internal/domain"
	"strconv"
	"strings"
)

func (s *Services) GetById(id int) (*domain.UserResponse, error) {
	user, err := s.Repo.GetLimitPedidosById(id)
	if err != nil {
		return nil, err
	}

	for _, pedido := range user.Pedidos {
		for i, producto := range pedido.Items {
			tasteIds := producto.Tastes
			ids := strings.Split(tasteIds, ",")
			var tasteIdsInt []int

			for _, id := range ids {
				idInt, err := strconv.Atoi(id)
				if err != nil {
					return nil, err
				}
				tasteIdsInt = append(tasteIdsInt, idInt)
			}

			tastesArray, err := s.RepoTastes.GetByIds(tasteIdsInt)
			if err != nil {
				return nil, err
			}

			pedido.Items[i].Tastes = ""

			for _, taste := range tastesArray {
				pedido.Items[i].Tastes += taste + ","
			}
		}
	}
	return user, nil
}
