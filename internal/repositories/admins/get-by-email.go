package admins

import (
	"clean_code/internal/domain"
)

func (r *AdminRepository) GetByEmail(email string) (*domain.Admin, error) {
	admin := &domain.Admin{}
	err := r.DB.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}
