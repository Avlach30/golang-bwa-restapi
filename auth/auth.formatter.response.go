package auth

type UserFormatterResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
}

type UserLoginFormatterResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUserSignupResponse(user User) UserFormatterResponse {
	format := UserFormatterResponse{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
	}

	return format
}

func FormatUserLoginResponse(user User, token string) UserLoginFormatterResponse {
	format := UserLoginFormatterResponse{
		ID:   user.ID,
		Email: user.Email,
		Token: token,
	}

	return format
}
