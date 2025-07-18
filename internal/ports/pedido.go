package ports

import "clean_code/internal/domain"

type PedidosRepository interface {
	Create(pedido *[]domain.Pedidos) (*[]domain.Pedidos, error)
	GetByUserId(userId int) ([]*domain.Pedidos, error)
	Update(comment *domain.Pedidos) (*domain.Pedidos, error)
	List() ([]*domain.Pedidos, error)
	Delete(id int) (*domain.Pedidos, error)
	GetById(id uint) (bool, error)
}

type PedidosService interface {
	Create(comment *domain.PedidosRequest) (*[]domain.Pedidos, error)
	GetByUserId(userId int) ([]*domain.Pedidos, error)
	Update(comment *domain.Pedidos) (*domain.Pedidos, error)
	List() ([]*domain.Pedidos, error)
	Delete(id int) (*domain.Pedidos, error)
}
