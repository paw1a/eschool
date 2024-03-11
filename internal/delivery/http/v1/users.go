package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/auth/sign-in", h.userSignIn)
		users.POST("/auth/logout", h.userLogout)
		users.POST("/auth/sign-up", h.userSignUp)
		users.POST("/auth/refresh", h.userRefresh)

		authenticated := users.Group("/", h.verifyUser)
		{
			authenticated.GET("/account", h.getUserAccount)
		}
	}
}

// UserAccount godoc
// @Summary   User account
// @Tags      user
// @Accept    json
// @Produce   json
// @Success   200  {object}  auth.AuthDetails
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  UserAuth
// @Router    /users/account [get]
func (h *Handler) getUserAccount(context *gin.Context) {
	userID, err := getIdFromRequestContext(context, "userID")
	if err != nil {
		errorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	userInfo, err := h.userService.FindUserInfo(context.Request.Context(), userID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, userInfo)
}
