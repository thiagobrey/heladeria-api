package handlers

import "clean_code/internal/ports"

type UserHandler struct {
	Services ports.UserService
}