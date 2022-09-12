package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SignUp(input SignUpInput) (User, error)
	LogIn(input LogInInput) (User, error)
	CheckUserAvailabilityByEmail(input SignUpInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

//* Mendeklarasikan fungsi untuk memproses input supaya bisa diproses oleh repository
func (service *service) SignUp(input SignUpInput) (User, error) {
	//* Create new user data
	user := User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email

	//* Hashing password
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(hashedPw)
	user.Role = "user"

	//* Memanggil method Save dari repository untuk pemroses hasil mapping instance struct User
	newUser, err := service.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (service *service) LogIn(input LogInInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	//* If user not found
	if user.ID == 0 {
		return user, errors.New("User with that email not found")
	}

	//* Compare hasing password and input password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	//* If error (compare result is different)
	if err != nil {
		return user, errors.New("incorrect Password")
	}

	return user, nil
}

func (service *service) CheckUserAvailabilityByEmail(input SignUpInput) (bool, error) {
	email := input.Email

	user, err := service.repository.FindByEmail(email)
	if (err != nil) {
		return false, nil
	}

	//* If user exist
	if (user.ID != 0) {
		return true, errors.New("User already exist")
	}

	return false, nil
}
