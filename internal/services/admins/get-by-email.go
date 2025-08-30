package admins

import "clean_code/internal/domain"

func (s *AdminServices) GetByEmail(email string) (*domain.Admin, error) {
	return s.RepoAdmin.GetByEmail(email)
}
