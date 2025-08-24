package admins

import "clean_code/internal/ports"

var _ ports.AdminService = &AdminServices{}

type AdminServices struct {
	RepoAdmin ports.AdminRepository
}
