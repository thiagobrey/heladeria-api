package repositories

import "clean_code/internal/domain"

func (r *Repository) Delete(id int) (*domain.User, error) {
	user := &domain.User{}
	err := r.DB.Delete(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
