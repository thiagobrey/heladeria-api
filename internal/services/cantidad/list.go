package cantidad

import "clean_code/internal/domain"

func (s *CantidadServices) List() ([]*domain.Cantidad, error) {
	cantidades, err := s.RepoCantidad.List()
	if err != nil {
		return nil, err
	}
	return cantidades, nil
}