package services

import "clean_code/internal/domain"

func (s *Services) GetUserByID(id int) (*domain.User, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
