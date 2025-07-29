package services

import (
	"clean_code/internal/ports"
)

var _ ports.PedidosService = &PedidosServices{}

type PedidosServices struct {
	RepoUser     ports.UserRepository
	RepoPedidos  ports.PedidosRepository
	RepoCantidad ports.CantidadRepository
	RepoTastes   ports.TasteRepository
}
