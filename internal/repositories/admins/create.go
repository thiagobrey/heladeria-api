package admins

import "clean_code/internal/domain"

func (r *AdminRepository) Create(admin *domain.Admin) (*domain.Admin, error) {
	err := r.DB.Create(&admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil

}
