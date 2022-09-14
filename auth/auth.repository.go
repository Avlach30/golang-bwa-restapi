package auth

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindUserById(id int) (User, error)
	UpdateTokenValue(token string, email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

//* Mendefinisikan function dengan mengembalikan nilai struct repository (db gorm)
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//* Mendefinisikan function/method untuk repository dengan create data dengan gorm
func (repo *repository) Save(user User) (User, error) {
	err := repo.db.Create(&user).Error

	if (err != nil) {
		return user, err
	} else {
		return user, nil
	}
}

func (repo *repository) FindByEmail(email string) (User, error) {
	var user User

	//* SELECT * FROM users WHERE email = ?
	err := repo.db.Where("email = ?", email).Find(&user).Error
	if (err != nil) {
		return user, err
	} else {
		return user, nil
	}
}

func (repo *repository) FindUserById(id int) (User, error) {
	var user User

	err := repo.db.Where("id = ?", id).Find(&user).Error
	if (err != nil) {
		return user, err
	} else {
		return user, nil
	}
}

func (repo *repository) UpdateTokenValue(token string, email string) (User, error) {
	var user User

	err := repo.db.Where("email = ?", email).Find(&user).Error
	if (err != nil) {
		return user, err
	} else {
		user.Token = token //* Update value
		repo.db.Save(&user) //* Save updated value

		return user, nil
	}
}