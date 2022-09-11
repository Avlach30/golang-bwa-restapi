package auth

import "golang.org/x/crypto/bcrypt"

type Service interface {
	SignUp(input SignUpInput) (User, error)
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
