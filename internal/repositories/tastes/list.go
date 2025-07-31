package tastes

import "clean_code/internal/domain"

func (r *TasteRepository) List() ([]*domain.Taste, error) {
	var tastes []*domain.Taste

	if err := r.DB.Model(&domain.Taste{}).Find(&tastes).Error; err != nil {
		return nil, err
	}

	return tastes, nil
}