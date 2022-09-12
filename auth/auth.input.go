package auth

type SignUpInput struct { 
	Name           string `form:"name" binding:"required"`
	Email          string `form:"email" binding:"required,email"`
	Occupation     string `form:"occupation" binding:"required"`
	Password       string `form:"password" binding:"required,min=8"`
	AvatarFileName string
}

type LogInInput struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}
