package tastes

import "clean_code/internal/domain"

func (s *TasteServices) List() ([]*domain.Taste, error) {
	tastes, err := s.Repo.List()
	if err != nil {
		return nil, err
	}
	return tastes, nil
}