package services

import "clean_code/internal/domain"

func (s *Services) Update(user *domain.UserRequest, id int) (*domain.User, error) {

	_, err := s.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	newUser := &domain.User{
		Id:        id,
		Name:      user.Name,
		Phone:     user.Phone,
		Direction: user.Direction,
	}

	updatedUser, err := s.Repo.Update(newUser)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
