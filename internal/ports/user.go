package ports

import "clean_code/internal/domain"

type UserRepository interface {
	GetUserByID(id int) (*domain.User, error)
	Delete(id int) (*domain.User, error)
	List() ([]*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	DirectionExists(direction string) (bool, error)
	Update(user *domain.User) (*domain.User, error)
}

type UserService interface {
	GetUserByID(id int) (*domain.User, error)
	Delete(id int) (*domain.User, error)
	List() ([]*domain.User, error)
	Create(user *domain.UserRequest) (*domain.User, error)
	Update(user *domain.UserRequest, id int) (*domain.User, error)
}
