package tastes

import "clean_code/internal/domain"

func (r *TasteRepository) Delete(id int) (*domain.Taste, error) {
	taste := &domain.Taste{}
	if err := r.DB.Delete(&taste, "id = ?" , id).Error; err != nil {
		return nil, err
	}
	return taste, nil
}
