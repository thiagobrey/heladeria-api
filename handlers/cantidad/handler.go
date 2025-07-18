package cantidad

import "clean_code/internal/ports"

type CantidadHandler struct {
	CantidadService ports.CantidadService
}