package ports

import "clean_code/internal/domain"

type TasteRepository interface {
	Create(taste *domain.Taste) (*domain.Taste, error)
	GetByIds(ids []int) ([]string, error)
}



type TasteService interface {
	Create(taste *domain.Taste) (*domain.Taste, error)
}
