package ports

import "clean_code/internal/domain"

type SesionsRepository interface {
	Create(sesion *domain.Sesions) (*domain.Sesions, error)
}

type SesionsService interface {
	Create(sesion *domain.Sesions) (*domain.Sesions, error)
}
