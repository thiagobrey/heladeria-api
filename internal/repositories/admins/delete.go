package admins

import "clean_code/internal/domain"

func (r *AdminRepository) Delete(id int) (*domain.Admin, error) {
	admin := &domain.Admin{}
	if err := r.DB.Delete(&admin, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return admin, nil
}
