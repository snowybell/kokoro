package repo

import "github.com/snowybell/kokoro/entity"

func (r *repository) GetToken(token entity.Token) (*entity.Token, error) {
	err := r.DB.Where(&token).Take(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *repository) SaveToken(token entity.Token) (*entity.Token, error) {
	err := r.DB.Save(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *repository) DeleteToken(token entity.Token) error {
	return r.DB.Delete(&token).Error
}
