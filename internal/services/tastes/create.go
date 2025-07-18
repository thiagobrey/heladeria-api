package tastes

import (
	"clean_code/internal/domain"
	"errors"
)

func (r *TasteServices) Create(taste *domain.Taste) (*domain.Taste, error) {
	if taste == nil {
		return nil, errors.New("taste cannot be nil")
	}

	if taste.Name == "" {
		return nil, errors.New("taste name is required")
	}

	newTaste, err := r.Repo.Create(taste)
	if err != nil {
		return nil, err
	}

	return newTaste, nil
}
