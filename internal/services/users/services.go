package services

import (
	"clean_code/internal/ports"

)

var _ ports.UserService = &Services{}

type Services struct {
	Repo ports.UserRepository
}
