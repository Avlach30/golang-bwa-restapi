package auth

type UserFormatterResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
}

func FormatUserSignupResponse(user User) UserFormatterResponse {
	format := UserFormatterResponse{
		ID: user.ID,
		Name: user.Name,
		Occupation: user.Occupation,
		Email: user.Email,
	}

	return format
}