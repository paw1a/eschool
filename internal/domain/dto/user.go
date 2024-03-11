package dto

type CreateUserDTO struct {
	Name     string `json:"name" required:"true"`
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

type UpdateUserDTO struct {
	Name string `json:"name"`
}

type UpdateUserInput struct {
	Name string `json:"name"`
}

type SignUpDTO struct {
	Name     string `json:"name" required:"true"`
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

type SignInDTO struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Fingerprint string `json:"fingerprint" binding:"required"`
}
