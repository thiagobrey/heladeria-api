package cantidad

import "clean_code/internal/domain"

func (s *CantidadServices) Delete(id int) (*domain.Cantidad, error) {
	cantidad, err := s.RepoCantidad.Delete(id)
	if err != nil {
		return nil, err
	}
	return cantidad, nil
}
