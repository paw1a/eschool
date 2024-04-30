package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
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

	userInfoDTO := dto.UserInfoDTO{
		Name:    userInfo.Name,
		Surname: userInfo.Surname,
	}

	SuccessResponse(context, userInfoDTO)
}

func (h *Handler) getAllUsers(context *gin.Context) {
	users, err := h.userService.FindAll(context.Request.Context())
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	userDTOs := make([]dto.UserDTO, len(users))
	if users != nil {
		for i, user := range users {
			userDTOs[i] = dto.UserDTO{
				ID:        user.ID.String(),
				Name:      user.Name,
				Surname:   user.Surname,
				Email:     user.Email,
				Phone:     user.Phone.String,
				City:      user.City.String,
				AvatarUrl: user.AvatarUrl.String,
			}
		}
	}

	SuccessResponse(context, userDTOs)
}
