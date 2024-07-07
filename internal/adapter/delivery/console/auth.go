package console

import (
	"context"
	"fmt"
	"github.com/paw1a/eschool/internal/adapter/delivery/console/dto"
	"github.com/paw1a/eschool/internal/core/port"
)

func (h *Handler) UserSignIn(c *Console) {
	var signInDTO dto.SignInDTO
	err := dto.InputSignInDTO(&signInDTO)
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println(signInDTO.Email)
	fmt.Println(signInDTO.Password)

	authDetails, err := h.authService.SignIn(context.Background(), port.SignInParam{
		Email:       signInDTO.Email,
		Password:    signInDTO.Password,
		Fingerprint: signInDTO.Fingerprint,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	user, err := h.userService.FindByCredentials(context.Background(), port.UserCredentials{
		Email:    signInDTO.Email,
		Password: signInDTO.Password,
	})

	c.UserID = &user.ID
	fmt.Printf("Access token: %s\n", authDetails.AccessToken.String())
}

func (h *Handler) UserSignUp(c *Console) {
	var signUpDTO dto.SignUpDTO
	err := dto.InputSignUpDTO(&signUpDTO)
	if err != nil {
		ErrorResponse(err)
		return
	}

	err = h.authService.SignUp(context.Background(), port.SignUpParam{
		Name:      signUpDTO.Name,
		Surname:   signUpDTO.Surname,
		Email:     signUpDTO.Email,
		Password:  signUpDTO.Password,
		Phone:     signUpDTO.Phone,
		City:      signUpDTO.City,
		AvatarUrl: signUpDTO.AvatarUrl,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("successfully signed up")
}

func (h *Handler) UserLogout(c *Console) {
	c.UserID = nil
	fmt.Println("successfully logged out")
}

func (h *Handler) verifyAuth(c *Console) error {
	if c.UserID == nil {
		return UnauthorizedError
	}
	return nil
}
