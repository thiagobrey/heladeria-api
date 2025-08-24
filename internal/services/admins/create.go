package admins

import "clean_code/internal/domain"

func (s *AdminServices) Create(admin *domain.Admin) (*domain.Admin, error) {
	admin, err := s.RepoAdmin.Create(admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}
