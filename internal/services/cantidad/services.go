package cantidad

import "clean_code/internal/ports"

var _ ports.CantidadService = &CantidadServices{}

type CantidadServices struct {
	RepoCantidad ports.CantidadRepository
}