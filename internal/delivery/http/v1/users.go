package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/domain"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		authenticated := users.Group("/", h.verifyToken)
		{
			users.GET("/", h.getAllUsers)
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

// GetUsers godoc
// @Summary   Get all users
// @Tags      user
// @Accept    json
// @Produce   json
// @Success   200  {array}   success
// @Failure   401  {object}  failure
// @Failure   404  {object}  failure
// @Failure   500  {object}  failure
// @Security  UserAuth
// @Router    /users [get]
func (h *Handler) getAllUsers(context *gin.Context) {
	users, err := h.userService.FindAll(context.Request.Context())
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	usersArray := make([]domain.User, len(users))
	if users != nil {
		usersArray = users
	}

	successResponse(context, usersArray)
}
