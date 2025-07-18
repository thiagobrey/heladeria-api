package repositories

import (
	"clean_code/internal/domain"
	"fmt"
)

func (r *Repository) Create(user *domain.User) (*domain.User, error) {
	err := r.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return user, nil
}
