package cantidad

import "clean_code/internal/domain"

func (s *CantidadServices) Create(cr *domain.CantidadRequest) (cantidad *domain.Cantidad, err error) {

	newCantidad := &domain.Cantidad{
		Description: cr.Description,
		CountTastes: cr.CountTastes,
		Code:        cr.Code,
	}

	cantidad, err = s.RepoCantidad.Create(newCantidad)
	if err != nil {
		return nil, err
	}

	return
}
