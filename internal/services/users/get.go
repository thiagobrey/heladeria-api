package services

import (
	"clean_code/internal/domain"
)

func (s *Services) GetById(id int) (*domain.UserResponse, error) {
	user, err := s.Repo.GetLimitPedidosById(id)
	if err != nil {
		return nil, err
	}

	return user, nil

}
