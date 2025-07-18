package cantidad

import "clean_code/internal/domain"

func (s *CantidadServices) Update(cr *domain.CantidadRequest, id int) (cantidad *domain.Cantidad, err error) {
	newCantidad := &domain.Cantidad{
		Id:          id,
		Description: cr.Description,
		CountTastes: cr.CountTastes,
		Code:        cr.Code,
	}

	cantidad, err = s.RepoCantidad.Update(newCantidad)
	if err != nil {
		return nil, err
	}
	return cantidad, nil
}
