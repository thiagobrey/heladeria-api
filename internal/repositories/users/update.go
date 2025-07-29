package repositories

import "clean_code/internal/domain"

func (r *Repository) Update(user *domain.User) (*domain.User, error) {
	err := r.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
