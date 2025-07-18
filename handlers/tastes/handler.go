package tastes

import "clean_code/internal/ports"

type TasteHandler struct {
	Services ports.TasteService
}