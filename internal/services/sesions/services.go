package sesions

import "clean_code/internal/ports"

var _ ports.SesionsService = &SesionsServices{}

type SesionsServices struct {
	Repo ports.SesionsRepository
}
