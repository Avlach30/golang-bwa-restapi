package auth

type SignUpInput struct {
	Name       string `json:"name" binding:"required"` //* binding property for handling validation
	Email      string `json:"email" binding:"required,email"`
	Occupation string `json:"occupation" binding:"required"`
	Password   string `json:"password" binding:"required,min=8"`
}
