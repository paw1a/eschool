package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/domain"
	"github.com/paw1a/eschool/internal/domain/dto"
	"github.com/paw1a/eschool/pkg/auth"
	"net/http"
)

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

	user, err := h.userService.FindByCredentials(context, signInDTO)
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

	user, err := h.userService.Create(context, dto.CreateUserDTO{
		Name:     signUpDTO.Name,
		Surname:  signUpDTO.Surname,
		Email:    signUpDTO.Email,
		Password: signUpDTO.Password,
	})
	if err != nil {
		internalErrorResponse(context, err)
		return
	}

	createdResponse(context, domain.UserInfo{
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
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

func (h *Handler) verifyUser(context *gin.Context) {
	h.verifyToken(context, "userID")
}
