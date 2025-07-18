package repositories

import "clean_code/internal/domain"

func (r *Repository) GetUserByID(id int) (*domain.User, error) {
	user := &domain.User{}
	err := r.DB.Preload("Comments").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
