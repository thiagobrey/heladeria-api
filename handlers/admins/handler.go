package admins

import "clean_code/internal/ports"

type AdminHandler struct {
	AdminService ports.AdminService
}
