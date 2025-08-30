package sesions

import (
	"clean_code/internal/domain"
	"errors"
)

func (s *SesionsServices) Create(sesion *domain.Sesions) (*domain.Sesions, error) {
	if sesion == nil {
		return nil, errors.New("sesion cannot be nil")
	}

	if sesion.Token == "" {
		return nil, errors.New("sesion token is required")
	}

	newSesions, err := s.Repo.Create(sesion)
	if err != nil {
		return nil, err
	}
	return newSesions, nil
}
