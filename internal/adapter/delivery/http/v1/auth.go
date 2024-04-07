package v1

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/paw1a/eschool/pkg/auth"
	log "github.com/sirupsen/logrus"
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

// UserSignIn godoc
// @Summary  User sign-in
// @Tags     user-auth
// @Accept   json
// @Produce  json
// @Param    user  body      dto.SignInDTO  true  "user credentials"
// @Success  200   {object}  auth.AuthDetails
// @Failure  400   {object}  failure
// @Failure  401   {object}  failure
// @Failure  404   {object}  failure
// @Failure  500   {object}  failure
// @Router   /users/auth/sign-in [post]
func (h *Handler) userSignIn(context *gin.Context) {
	var signInDTO dto.SignInDTO
	err := context.BindJSON(&signInDTO)
	if err != nil {
		badRequestResponse(context, "invalid sign in credentials", err)
		return
	}

	user, err := h.userService.FindByCredentials(context, port.UserCredentials{
		Email:    signInDTO.Email,
		Password: signInDTO.Password,
	})
	if err != nil {
		unauthorizedResponse(context, "invalid user email or password")
		return
	}

	userClaims := jwt.MapClaims{"userID": user.ID}
	authDetails, err := h.tokenProvider.CreateJWTSession(auth.CreateSessionInput{
		Fingerprint: signInDTO.Fingerprint,
		Claims:      userClaims,
	})

	if err != nil {
		internalErrorResponse(context, err)
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("refreshToken", authDetails.RefreshToken,
		int(h.config.JWT.RefreshTokenTime), "/", h.config.Server.Host, false, false)

	successResponse(context, authDetails.AccessToken)
}

// UserSignUp godoc
// @Summary  User sign-up
// @Tags     user-auth
// @Accept   json
// @Produce  json
// @Param    user  body      dto.SignUpDTO  true  "user data"
// @Success  200   {object}  domain.UserInfo
// @Failure  400   {object}  failure
// @Failure  401   {object}  failure
// @Failure  404   {object}  failure
// @Failure  500   {object}  failure
// @Router   /users/auth/sign-up [post]
func (h *Handler) userSignUp(context *gin.Context) {
	var signUpDTO dto.SignUpDTO
	err := context.BindJSON(&signUpDTO)
	if err != nil {
		badRequestResponse(context, "invalid sign up data", err)
		return
	}

	user, err := h.userService.Create(context, port.CreateUserParam{
		Name:      signUpDTO.Name,
		Surname:   signUpDTO.Surname,
		Email:     signUpDTO.Email,
		Password:  signUpDTO.Password,
		Phone:     signUpDTO.Phone,
		City:      signUpDTO.City,
		AvatarUrl: signUpDTO.AvatarUrl,
	})
	if err != nil {
		internalErrorResponse(context, err)
		return
	}

	createdResponse(context, dto.UserInfo{
		Name:    user.Name,
		Surname: user.Surname,
	})
}

// UserRefresh godoc
// @Summary  User refresh token
// @Tags     user-auth
// @Accept   json
// @Produce  json
// @Param    refreshInput  body      auth.RefreshInput  true  "user refresh data"
// @Success  200           {object}  auth.AuthDetails
// @Failure  400           {object}  failure
// @Failure  401           {object}  failure
// @Failure  404           {object}  failure
// @Failure  500           {object}  failure
// @Router   /users/auth/refresh [post]
func (h *Handler) userRefresh(context *gin.Context) {
	h.refreshToken(context)
}

// UserLogout godoc
// @Summary  User logout
// @Tags     user-auth
// @Accept   json
// @Produce  json
// @Failure  400           {object}  failure
// @Failure  401           {object}  failure
// @Failure  404           {object}  failure
// @Failure  500           {object}  failure
// @Router   /users/auth/logout [post]
func (h *Handler) userLogout(context *gin.Context) {
	refreshCookie, err := context.Cookie("refreshToken")
	if err != nil {
		unauthorizedResponse(context, "you are already logout")
		return
	}

	err = h.tokenProvider.DeleteJWTSession(refreshCookie)
	if err != nil {
		internalErrorResponse(context, err)
		return
	}

	context.SetCookie("refreshToken", "", -1, "/", h.config.Server.Host, false, false)

	successResponse(context, nil)
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

	log.Printf("token = %s", headerParts[1])

	return headerParts[1], nil
}

func (h *Handler) refreshToken(context *gin.Context) {
	var input auth.RefreshInput

	err := context.BindJSON(&input)
	if err != nil {
		badRequestResponse(context, "can't parse request body", err)
		return
	}

	refreshCookie, err := context.Cookie("refreshToken")
	if err != nil {
		badRequestResponse(context, "refresh cookie not found", err)
		return
	}

	input.RefreshToken = refreshCookie

	authDetails, err := h.tokenProvider.Refresh(auth.RefreshInput{
		RefreshToken: input.RefreshToken,
		Fingerprint:  input.Fingerprint,
	})

	if err != nil {
		unauthorizedResponse(context, err.Error())
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("refreshToken", authDetails.RefreshToken,
		int(h.config.JWT.RefreshTokenTime), "/", h.config.Server.Host, false, false)

	successResponse(context, authDetails.AccessToken)
}

func (h *Handler) verifyToken(context *gin.Context) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		unauthorizedResponse(context, err.Error())
		return
	}

	tokenClaims, err := h.tokenProvider.VerifyToken(tokenString)
	if err != nil {
		unauthorizedResponse(context, err.Error())
		return
	}

	id, ok := tokenClaims["userID"]
	if !ok {
		errorResponse(context, http.StatusForbidden, "this endpoint is forbidden")
		return
	}

	context.Set("userID", id)
}

func (h *Handler) extractIdFromAuthHeader(context *gin.Context, idName string) (domain.ID, error) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		return domain.RandomID(), err
	}

	tokenClaims, err := h.tokenProvider.VerifyToken(tokenString)
	if err != nil {
		return domain.RandomID(), err
	}

	id, ok := tokenClaims[idName]
	if !ok {
		return domain.RandomID(), fmt.Errorf("failed to extract %s from auth header", idName)
	}

	return id.(domain.ID), nil
}
