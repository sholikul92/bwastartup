package user

import "gorm.io/gorm"

type UserRepository interface {
	Save(user Users) (*Users, error)
	FindByEmail(email string) (Users, error)
	FindByID(ID int) (Users, error)
	Update(user Users) (Users, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Save(user Users) (*Users, error) {

	errSave := r.db.Create(&user).Error
	if errSave != nil {
		return nil, errSave
	}

	return &user, nil
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

func (r *Repository) FindByID(ID int) (Users, error) {
	var user Users

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) Update(user Users) (Users, error) {

	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
