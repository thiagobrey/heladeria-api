package tastes

import "clean_code/internal/domain"

func (s *TasteServices) GetById(id int) (*domain.Taste, error) {
	taste, err := s.Repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return taste, nil
}