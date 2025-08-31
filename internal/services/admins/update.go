package admins

import "clean_code/internal/domain"

func (s *AdminServices) Update(admin *domain.Admin) (*domain.Admin, error) {
	updatedAdmin, err := s.RepoAdmin.Update(admin)
	if err != nil {
		return nil, err
	}
	return updatedAdmin, nil
}
