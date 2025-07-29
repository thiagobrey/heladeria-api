package cantidad

import "clean_code/internal/domain"

func (r *CantidadRepository) List() ([]*domain.Cantidad, error) {
	var cantidades []*domain.Cantidad

	if err := r.DB.Model(&domain.Cantidad{}).Find(&cantidades).Error; err != nil {
		return nil, err
	}

	return cantidades, nil
}
