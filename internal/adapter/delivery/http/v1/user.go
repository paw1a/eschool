package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	"github.com/paw1a/eschool/internal/core/port"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.GET("/", h.findAllUsers)
		authenticated := users.Group("/", h.verifyToken)
		{
			authenticated.GET("/account", h.findUserAccount)
			authenticated.PUT("/account", h.updateUser)

			authenticated.GET("/courses", h.findUserCourses)
			authenticated.GET("/certificates", h.findUserCertificates)
		}
	}
}

func (h *Handler) findAllUsers(context *gin.Context) {
	users, err := h.userService.FindAll(context.Request.Context())
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	userDTOs := make([]dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.NewUserDTO(user)
	}

	SuccessResponse(context, userDTOs)
}

func (h *Handler) findUserAccount(context *gin.Context) {
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

func (h *Handler) updateUser(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		ErrorResponse(context, UnauthorizedError)
		return
	}

	var updateUserDTO dto.UpdateUserDTO
	err = context.ShouldBindJSON(&updateUserDTO)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	user, err := h.userService.Update(context.Request.Context(), userID, port.UpdateUserParam{
		Name:      updateUserDTO.Name,
		Surname:   updateUserDTO.Surname,
		Phone:     updateUserDTO.Phone,
		City:      updateUserDTO.City,
		AvatarUrl: updateUserDTO.AvatarUrl,
	})

	userDTO := dto.NewUserDTO(user)
	SuccessResponse(context, userDTO)
}

func (h *Handler) findUserCourses(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		ErrorResponse(context, UnauthorizedError)
		return
	}

	courses, err := h.courseService.FindStudentCourses(context.Request.Context(), userID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	courseDTOs := make([]dto.CourseDTO, len(courses))
	for i, course := range courses {
		courseDTOs[i] = dto.NewCourseDTO(course)
	}

	SuccessResponse(context, courseDTOs)
}

func (h *Handler) findUserCertificates(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		ErrorResponse(context, UnauthorizedError)
		return
	}

	courses, err := h.courseService.FindStudentCourses(context.Request.Context(), userID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	courseDTOs := make([]dto.CourseDTO, len(courses))
	for i, course := range courses {
		courseDTOs[i] = dto.NewCourseDTO(course)
	}

	SuccessResponse(context, courseDTOs)
}
