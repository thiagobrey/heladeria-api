package repositories

import (
	"clean_code/internal/domain"
	"errors"

	"gorm.io/gorm"
)

func (r *Repository) DirectionExists(direction string) (bool, error) {
	user := &domain.User{}
	err := r.DB.First(&user, "direction = ?", direction).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}

	if user.Id > 0 {
		return false, errors.New("direction already exists")
	}

	return true, nil
}
