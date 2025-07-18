package cantidad

import (
	"clean_code/internal/domain"
	
)

func (r *CantidadRepository) Create(cantidad *domain.Cantidad) (*domain.Cantidad, error) {
	err := r.DB.Create(&cantidad).Error
	if err != nil {
		return nil, err
	}

	return cantidad, nil
}
