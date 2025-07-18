package services

import (
	"clean_code/internal/domain"
)

func (s *Services) Create(ur *domain.UserRequest) (user *domain.User, err error) {
	_, err1 := s.Repo.DirectionExists(ur.Direction)
	if err1 != nil {
		return nil, err1
	}

	newUser := &domain.User{
		Name:      ur.Name,
		Phone:     ur.Phone,
		Direction: ur.Direction,
	}

	user, err = s.Repo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return
}
