package admins

import "clean_code/internal/domain"

func (r *AdminRepository) GetById(id int) (*domain.Admin, error) {
	var admin domain.Admin
	err := r.DB.First(&admin, id).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
