package ports

import "clean_code/internal/domain"

type CantidadRepository interface {
	Create(cr *domain.Cantidad) (*domain.Cantidad, error)
	Update(cr *domain.Cantidad) (cantidad *domain.Cantidad, err error)
	GetByCode(code int) (*domain.Cantidad, error)
	List() ([]*domain.Cantidad, error)
}

type CantidadService interface {
	Create(cr *domain.CantidadRequest) (*domain.Cantidad, error)
	Update(cr *domain.CantidadRequest, id int) (cantidad *domain.Cantidad, err error)
	List() ([]*domain.Cantidad, error)
}
