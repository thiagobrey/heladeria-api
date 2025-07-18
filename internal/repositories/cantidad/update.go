package cantidad

import "clean_code/internal/domain"

func (r *CantidadRepository) Update(cantidad *domain.Cantidad) (*domain.Cantidad, error) {
	err := r.DB.Save(&cantidad).Error
	if err != nil {
		return nil, err
	}
	return cantidad, nil
}
