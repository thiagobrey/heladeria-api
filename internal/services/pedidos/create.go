package services

import (
	"clean_code/internal/domain"
	"strconv"
	"strings"
)

func (s *PedidosServices) Create(request *domain.PedidosRequest) (*domain.PedidosResponse, error) {

	var Pedidos []domain.Pedidos
	var cantidad *domain.Cantidad
	var total float32
	var items []domain.Productos

	user, err := s.RepoUser.GetById(request.UserId)
	if err != nil {
		return nil, err
	}

	for _, producto := range request.Productos {
		cantidad, err = s.RepoCantidad.GetByCode(producto.Code)
		if err != nil {
			return nil, err
		}

		Pedidos = append(Pedidos, domain.Pedidos{
			UserId: request.UserId,
			Price:  float32(cantidad.Price),
			Code:   cantidad.Code,
			Tastes: producto.Tastes,
		})
		total += float32(cantidad.Price)
	}

	_, err = s.RepoPedidos.Create(&Pedidos)
	if err != nil {
		return nil, err
	}

	// Slice de gustos (check)
	// Con los id de los gustos, hacer un get a la tabla de gustos con un where in (arreglar)
	// Obtener el nombre de cada gusto y guardarlo en el PedidoResponse
	for _, pedido := range Pedidos {
		tasteIds := pedido.Tastes
		ids := strings.Split(tasteIds, ",")
		var tasteIdsInt []int
		for _, id := range ids {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return nil, err
			}
			tasteIdsInt = append(tasteIdsInt, idInt)
		}
		tastes, err := s.RepoTastes.GetByIds(tasteIdsInt)
		if err != nil {
			return nil, err
		}
		items = append(items, domain.Productos{
			Code:   pedido.Code,
			Price:  pedido.Price,
			Tastes: strings.Join(tastes, ","),
		})
	}

	response := &domain.PedidosResponse{
		User:  *user,
		Items: items,
		Total: total,
	}

	return response, nil
}
