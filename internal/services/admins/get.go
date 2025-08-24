package admins

import "clean_code/internal/domain"

func (s *AdminServices) GetById(id int) (*domain.Admin, error) {
	admin, err := s.RepoAdmin.GetById(id)
	if err != nil {
		return nil, err
	}
	return admin, nil
}
