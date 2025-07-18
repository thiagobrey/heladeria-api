package services

import (
	"clean_code/internal/domain"
)

func (s *PedidosServices) Create(request *domain.PedidosRequest) (*[]domain.Pedidos, error) {

	var Pedidos []domain.Pedidos
	var cantidad domain.Cantidad

	_, err := s.RepoUser.GetUserByID(request.UserId)
	if err != nil {
		return nil, err
	}

	for _, producto := range request.Productos {
		// TODO: Obtener el precio de la cantidad por producto.Code
		switch cantidad.Code {
		case 104:
			producto.Price = 1000.00
		case 102:
			producto.Price = 1800.00
		case 103:
			producto.Price = 2600.00
		case 101:
			producto.Price = 3500.00
		}
		Pedidos = append(Pedidos, domain.Pedidos{
			UserId:   request.UserId,
			Price:    producto.Price,
			Cantidad: producto.Cantidad,
			Tastes:   producto.Tastes,
		})
	}

	newPedido, err := s.RepoPedidos.Create(&Pedidos)
	if err != nil {
		return nil, err
	}

	return newPedido, nil
}
