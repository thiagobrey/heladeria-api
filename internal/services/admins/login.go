package admins

import (
	"clean_code/internal/domain"
	"clean_code/middleware"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *AdminServices) Login(email, password string) (string, error) {
	admin, err := s.RepoAdmin.GetByEmail(email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)) != nil {
		return "", errors.New("Credenciales inválidas")
	}

	token, err := middleware.GenerateToken(admin.Email)
	if err != nil {
		return "", errors.New("No se pudo generar el token")
	}

	sesion := domain.Sesions{
		Token:     token,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Data:      admin.Email,
	}
	_, err = s.RepoSesion.Create(&sesion)
	if err != nil {
		return "", errors.New("No se pudo crear la sesión")
	}

	return token, nil
}
