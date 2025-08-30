package sesions

import "clean_code/internal/ports"

type SesionHandler struct {
	SesionService ports.SesionsService
}