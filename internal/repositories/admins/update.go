package admins

import "clean_code/internal/domain"

func (r *AdminRepository) Update(admin *domain.Admin) (*domain.Admin, error) {
	err := r.DB.Save(&admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}
