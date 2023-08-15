package user

import "gorm.io/gorm"

type UserRepository interface {
	Save(user Users) (Users, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Save(user Users) (Users, error) {

	errSave := r.db.Save(&user).Error
	if errSave != nil {
		return user, errSave
	}

	return user, nil
}
