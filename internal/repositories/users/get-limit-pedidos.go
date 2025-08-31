package repositories

import (
	"clean_code/internal/domain"
	"strings"

	"gorm.io/gorm"
)

func (r *Repository) GetLimitPedidosById(id int) (*domain.UserResponse, error) {
	var rawFechas []string
	err := r.DB.
		Model(&domain.Pedidos{}).
		Select("DATE(created_at)").
		Where("user_id = ?", id).
		Group("DATE(created_at)").
		Order("DATE(created_at) desc").
		Limit(3).
		Pluck("DATE(created_at)", &rawFechas).Error
	if err != nil {
		return nil, err
	}

	// Convertir las fechas al formato correcto (YYYY-MM-DD)
	var fechas []string
	for _, rawFecha := range rawFechas {
		// Eliminar cualquier formato ISO 8601 adicional
		fecha := strings.Split(rawFecha, "T")[0]
		fechas = append(fechas, fecha)
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

	// Agrupar pedidos por días
	diasMap := make(map[string][]domain.Items)
	for _, pedido := range user.Pedidos {
		createdAt := pedido.CreatedAt.Format("2006-01-02") // Formatear la fecha
		item := domain.Items{
			Price:  pedido.Price,
			Code:   pedido.Code,
			Tastes: pedido.Tastes,
		}
		diasMap[createdAt] = append(diasMap[createdAt], item)
	}

	// Construir la respuesta
	var dias []domain.PedidosUser
	for fecha, items := range diasMap {
		dias = append(dias, domain.PedidosUser{
			Fecha: fecha,
			Items: items,
		})
	}

	user.Pedidos = nil

	response := &domain.UserResponse{
		User:    *user,
		Pedidos: dias,
	}

	return response, nil
}
