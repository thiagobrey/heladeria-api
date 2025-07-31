package cantidad

import "clean_code/internal/domain"

func (r *CantidadRepository) GetById(id int) (*domain.Cantidad, error) {
	cantidad := &domain.Cantidad{}
	err := r.DB.Model(&domain.User{}).First(&cantidad, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return cantidad, nil
}