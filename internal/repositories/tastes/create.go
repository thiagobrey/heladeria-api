package tastes

import "clean_code/internal/domain"

func (r *TasteRepository) Create(taste *domain.Taste) (*domain.Taste, error) {
	if err := r.DB.Create(taste).Error; err != nil {
		return nil, err
	}
	return taste, nil
}
