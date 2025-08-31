package ports

import "clean_code/internal/domain"

type TasteRepository interface {
	Create(taste *domain.Taste) (*domain.Taste, error)
	GetByIds(ids []int) ([]string, error)
	Delete(id int) (*domain.Taste, error)
	GetById(id int) (*domain.Taste, error)
	Update(taste *domain.Taste) (*domain.Taste, error)
	List() ([]*domain.Taste, error) 
}



type TasteService interface {
	Create(taste *domain.Taste) (*domain.Taste, error)
	Delete(id int) (*domain.Taste, error)
	GetById(id int) (*domain.Taste, error)
	Update(taste *domain.Taste) (*domain.Taste, error)
	List() ([]*domain.Taste, error) 
}
