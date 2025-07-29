package cantidad

import (
	"clean_code/internal/domain"
	"fmt"
)

func (r *CantidadRepository) GetByCode(code int) (*domain.Cantidad, error) {
	var cantidad domain.Cantidad

	if err := r.DB.Model(&domain.Cantidad{}).Where("code = ?", code).Debug().First(&cantidad).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &cantidad, nil

}
