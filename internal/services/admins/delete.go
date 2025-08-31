package admins

import "clean_code/internal/domain"

func (s *AdminServices) Delete(id int) (*domain.Admin, error) {
	admin, err := s.RepoAdmin.Delete(id)
	if err != nil {
		return nil, err
	}
	return admin, nil
}
