package ports

import "clean_code/internal/domain"

type AdminRepository interface {
	Create(admin *domain.Admin) (*domain.Admin, error)
	GetById(id int) (*domain.Admin, error)
	Update(admin *domain.Admin) (*domain.Admin, error)
	Delete(id int) (*domain.Admin, error)
}

type AdminService interface {
	Create(admin *domain.Admin) (*domain.Admin, error)
	GetById(id int) (*domain.Admin, error)
	Update(admin *domain.Admin) (*domain.Admin, error)
	Delete(id int) (*domain.Admin, error)
}
