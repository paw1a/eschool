package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"net/http"
	"strings"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/sign-in", h.userSignIn)
		authGroup.POST("/logout", h.userLogout)
		authGroup.POST("/sign-up", h.userSignUp)
		authGroup.POST("/refresh", h.userRefresh)
	}
}

func (h *Handler) userSignIn(context *gin.Context) {
	var signInDTO dto.SignInDTO
	err := context.ShouldBindJSON(&signInDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	authDetails, err := h.authService.SignIn(context, port.SignInParam{
		Email:       signInDTO.Email,
		Password:    signInDTO.Password,
		Fingerprint: signInDTO.Fingerprint,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("refreshToken", authDetails.RefreshToken.String(),
		86400, "/", h.config.Host, false, false)

	h.successResponse(context, authDetails.AccessToken.String())
}

func (h *Handler) userSignUp(context *gin.Context) {
	var signUpDTO dto.SignUpDTO
	err := context.ShouldBindJSON(&signUpDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.authService.SignUp(context, port.SignUpParam{
		Name:      signUpDTO.Name,
		Surname:   signUpDTO.Surname,
		Email:     signUpDTO.Email,
		Password:  signUpDTO.Password,
		Phone:     signUpDTO.Phone,
		City:      signUpDTO.City,
		AvatarUrl: signUpDTO.AvatarUrl,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.createdResponse(context, "successfully signed up")
}

func (h *Handler) userRefresh(context *gin.Context) {
	h.refreshToken(context)
}

func (h *Handler) userLogout(context *gin.Context) {
	refreshCookie, err := context.Cookie("refreshToken")
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	err = h.authService.LogOut(context, domain.Token(refreshCookie))
	if err != nil {
		h.errorResponse(context, err)
	}

	context.SetCookie("refreshToken", "", -1, "/", h.config.Host, false, false)

	h.successResponse(context, "successfully logged out")
}

func (h *Handler) refreshToken(context *gin.Context) {
	var refreshDTO dto.RefreshDTO
	err := context.ShouldBindJSON(&refreshDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	refreshCookie, err := context.Cookie("refreshToken")
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	authDetails, err := h.authService.Refresh(context, domain.Token(refreshCookie),
		refreshDTO.Fingerprint)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("refreshToken", authDetails.RefreshToken.String(),
		86400, "/", h.config.Host, false, false)

	h.successResponse(context, authDetails.AccessToken.String())
}

func (h *Handler) verifyToken(context *gin.Context) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	payload, err := h.authService.Payload(context, domain.Token(tokenString))
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	context.Set("userID", payload.UserID.String())
}

func extractAuthToken(context *gin.Context) (string, error) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}

func (h *Handler) extractIdFromAuthHeader(context *gin.Context) (domain.ID, error) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		return domain.RandomID(), err
	}

	payload, err := h.authService.Payload(context, domain.Token(tokenString))
	if err != nil {
		h.errorResponse(context, err)
		return domain.RandomID(), err
	}

	return payload.UserID, nil
}
