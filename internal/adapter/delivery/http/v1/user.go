package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/core/domain"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.GET("/", h.getAllUsers)
		authenticated := users.Group("/", h.verifyToken)
		{
			authenticated.GET("/account", h.getUserAccount)
		}
	}
}

func (h *Handler) getUserAccount(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		ErrorResponse(context, UnauthorizedError)
		return
	}

	userInfo, err := h.userService.FindUserInfo(context.Request.Context(), userID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	SuccessResponse(context, userInfo)
}

func (h *Handler) getAllUsers(context *gin.Context) {
	users, err := h.userService.FindAll(context.Request.Context())
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	usersArray := make([]domain.User, len(users))
	if users != nil {
		usersArray = users
	}

	SuccessResponse(context, usersArray)
}
