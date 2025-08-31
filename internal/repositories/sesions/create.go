package sesions

import "clean_code/internal/domain"

func (r *SesionsRepository) Create(sesion *domain.Sesions) (*domain.Sesions, error) {
	err := r.DB.Create(&sesion).Error
	if err != nil {
		return nil, err
	}

	return sesion, nil
}
