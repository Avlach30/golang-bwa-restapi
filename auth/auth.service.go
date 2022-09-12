package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SignUp(input SignUpInput) (User, error)
	LogIn(input LogInInput) (User, string, error)
	CheckUserAvailabilityByEmail(input SignUpInput) (bool, error)
	GenerateToken(userId int, email string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (service *service) LogIn(input LogInInput) (User, string, error) {
	email := input.Email
	password := input.Password

	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return user, "", err
	}

	//* If user not found
	if user.ID == 0 {
		return user, "", errors.New("User with that email not found")
	}

	//* Compare hasing password and input password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	//* If error (compare result is different)
	if err != nil {
		return user, "", errors.New("incorrect Password")
	}

	userToken, err := service.GenerateToken(user.ID, user.Email)
	if err != nil {
		return user, "", errors.New("failed to generate token")
	}

	//* Call repository for update token
	service.repository.UpdateTokenValue(userToken, email)

	return user, userToken, nil
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

func (service *service) GenerateToken(userId int, email string) (string, error) {

	//* Generate payload token
	claim := jwt.MapClaims{
		"userId": userId,
		"email": email,
	}

	//* Generate token with signing method and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//* Signature token with secret text
	signedToken, err := token.SignedString([]byte("$2a$08$UcyjEygcPA/XaeUp85sQjuOhithx14/Ai3D5lYPixLrMrSQG2NIFy"))
	if (err != nil) {
		return signedToken, err
	}

	return signedToken, nil
}

func (service *service) ValidateToken(encodedToken string) (*jwt.Token, error) {

	//* Parse jwt for get token with valid sign secret text
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token)(interface{}, error){
		//* Check token signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if (!ok) {
			return nil, errors.New("invalid token")
		}

		return []byte("$2a$08$UcyjEygcPA/XaeUp85sQjuOhithx14/Ai3D5lYPixLrMrSQG2NIFy"), nil
	})

	if (err != nil) {
		return token, err
	}

	return token, nil

}
