package repositories

import (
	"clean_code/internal/domain"

	"gorm.io/gorm"
)

func (r *Repository) GetLimitPedidosById(id int) (*domain.User, error) {
	var fechas []string
	err := r.DB.
		Model(&domain.Pedidos{}).
		Select("DATE(created_at)").
		Where("user_id = ?", id).
		Group("DATE(created_at)").
		Order("DATE(created_at) desc").
		Limit(3).
		Pluck("DATE(created_at)", &fechas).Error
	if err != nil {
		return nil, err
	}


	user := &domain.User{}
	err = r.DB.Model(&domain.User{}).
		Preload("Pedidos", func(db *gorm.DB) *gorm.DB {
			return db.Where("DATE(created_at) IN ?", fechas).Order("created_at desc")
		}).
		First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
