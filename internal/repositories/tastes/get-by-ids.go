package tastes

import "clean_code/internal/domain"

/*
*retorna los gustos asociados a los ids proporcionado
 */
func (r *TasteRepository) GetByIds(ids []int) ([]string, error) {
	var tastes []string
	if err := r.DB.Model(&domain.Taste{}).Where("id IN ?", ids).Debug().Pluck("name", &tastes).Error; err != nil {
		return nil, err
	}
	return tastes, nil
}
