package tastes

import "clean_code/internal/domain"

func (r *TasteRepository) GetById(id int) (*domain.Taste, error) {
	taste := &domain.Taste{}
	err := r.DB.Model(&domain.Taste{}).First(&taste, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return taste, nil
}