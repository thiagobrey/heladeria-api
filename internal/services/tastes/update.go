package tastes

import "clean_code/internal/domain"

func (s *TasteServices) Update(taste *domain.Taste) (*domain.Taste, error) {
	updatedTaste, err := s.Repo.Update(taste)
	if err != nil {
		return nil, err
	}

	return updatedTaste, nil
}
