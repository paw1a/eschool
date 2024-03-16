package dto

type CreateUserDTO struct {
	Name     string `json:"name" required:"true"`
	Surname  string `json:"surname" required:"true"`
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

type UpdateUserDTO struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type UpdateUserInput struct {
	Name string `json:"name"`
}

type SignUpDTO struct {
	Name     string `json:"name" required:"true"`
	Surname  string `json:"surname" required:"true"`
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

type SignInDTO struct {
	Email       string `json:"email" required:"true"`
	Password    string `json:"password" required:"true"`
	Fingerprint string `json:"fingerprint" required:"true"`
}
