package services

import "clean_code/internal/domain"

func (s *Services) Delete(id int) (*domain.User, error) {

	_, err := s.Repo.GetById(id)
	if err != nil {
		return nil, err
	}

	user, err := s.Repo.Delete(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
