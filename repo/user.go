package repo

import "github.com/snowybell/kokoro/entity"

func (r *repository) GetUser(user entity.User) (*entity.User, error) {
	err := r.DB.Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) SaveUser(user entity.User) error {
	return r.DB.Save(&user).Error
}

func (r *repository) DeleteUser(user entity.User) error {
	return r.DB.Delete(&user).Error
}
