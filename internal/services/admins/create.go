package admins

import (
	"clean_code/hashpassword"
	"clean_code/internal/domain"
	"errors"
)

func (s *AdminServices) Create(admin *domain.Admin) (*domain.Admin, error) {
	hashed, err := hashpassword.PasswordHashing(admin.Password)
	if err != nil {
		return nil, err
	}
	admin.Password = hashed

	newAdmin, err := s.RepoAdmin.GetByEmail(admin.Email)
	if err == nil && newAdmin != nil {
		return nil, errors.New("El email ya está registrado")
	}

	createdAdmin, err := s.RepoAdmin.Create(admin)
	if err != nil {
		return nil, err
	}
	return createdAdmin, nil
}
