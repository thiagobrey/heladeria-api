package tastes

import (
	"clean_code/internal/ports"
)

var _ ports.TasteService = &TasteServices{}

type TasteServices struct {
	Repo ports.TasteRepository
}