package services

import (
	"clean_code/internal/ports"

)

var _ ports.UserService = &Services{}

type Services struct {
	Repo ports.UserRepository
	RepoCantidad ports.CantidadRepository
	RepoTastes ports.TasteRepository
	RepoPedidos ports.PedidosRepository
}
