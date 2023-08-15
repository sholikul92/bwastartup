package user

import "gorm.io/gorm"

type UserRepository interface {
	Save(user Users) (Users, error)
	FindByEmail(email string) (Users, error)
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

// Mencari email didalam database
func (r *Repository) FindByEmail(email string) (Users, error) {
	var user Users

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
