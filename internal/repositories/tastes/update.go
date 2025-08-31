package tastes

import "clean_code/internal/domain"

func (r *TasteRepository) Update(taste *domain.Taste) (*domain.Taste, error) {
	err := r.DB.Save(&taste).Error
	if err != nil {
		return nil, err
	}
	return taste, nil
}