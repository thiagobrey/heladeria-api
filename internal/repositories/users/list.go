package repositories

import "clean_code/internal/domain"

func (r *Repository) List() ([]*domain.User, error) {
	users := []*domain.User{}
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
