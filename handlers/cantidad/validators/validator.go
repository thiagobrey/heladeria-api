package validators

import (
	"clean_code/internal/domain"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CountValidator(fl validator.FieldLevel, db *gorm.DB) bool {
	

	producto, ok := fl.Field().Interface().(domain.Productos)
	if !ok {
		return false
	}
	
	var cantidad domain.Cantidad
	if err := db.Where("code = ?", producto.Code).First(&cantidad).Error; err != nil {
	
		return false
	}

	tasteIds := producto.Tastes
	ids := strings.Split(tasteIds, ",")
	
	if cantidad.CountTastes < len(ids) {
		return false
	}
	//}

	return true
}
