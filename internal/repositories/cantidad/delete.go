package cantidad

import "clean_code/internal/domain"

func (r *CantidadRepository) Delete(id int) (*domain.Cantidad, error) {
	cantidad := &domain.Cantidad{}
	if err := r.DB.Delete(&cantidad, "id = ?" , id).Error; err != nil {
		return nil, err
	}
	return cantidad, nil
}