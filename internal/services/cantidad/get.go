package cantidad

import "clean_code/internal/domain"

func (s *CantidadServices) GetById(id int) (*domain.Cantidad, error) {
	cantidad, err := s.RepoCantidad.GetById(id)
	if err != nil {
		return nil, err
	}
	return cantidad, nil
}
