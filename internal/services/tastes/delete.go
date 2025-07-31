package tastes

import "clean_code/internal/domain"

func (s *TasteServices) Delete(id int) (*domain.Taste, error) {
	taste, err := s.Repo.Delete(id)
	if err != nil {
		return nil, err
	}
	return taste, nil
}
