package services

import "clean_code/internal/domain"

func (s *Services) List() ([]*domain.User, error) {
	users, err := s.Repo.List()
	if err != nil {
		return nil, err
	}
	return users, nil
}
