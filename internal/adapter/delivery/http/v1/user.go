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
		users.GET("/:id", h.findUserByID)
		authenticated := users.Group("/", h.verifyToken)
		{
			authenticated.GET("/me", h.findUserAccount)
			authenticated.PATCH("/me", h.updateUser)

			authenticated.GET("/me/courses", h.findUserCourses)
			authenticated.PUT("/me/courses/:course_id", h.addUserFreeCourse)
		}
	}
}

// @Summary GetAllUsers
// @Tags user
// @Description get all users
// @Accept  json
// @Produce json
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} []dto.UserDTO
// @Router /users [get]
func (h *Handler) findAllUsers(context *gin.Context) {
	users, err := h.userService.FindAll(context.Request.Context())
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	userDTOs := make([]dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.NewUserDTO(user)
	}

	h.successResponse(context, userDTOs)
}

// @Summary GetUserByID
// @Tags user
// @Description get user by id
// @Accept  json
// @Produce json
// @Param   id   path    string  true  "user id"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} []dto.UserDTO
// @Router /users/{id} [get]
func (h *Handler) findUserByID(context *gin.Context) {
	userID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	user, err := h.userService.FindByID(context.Request.Context(), userID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	userDTO := dto.NewUserDTO(user)
	h.successResponse(context, userDTO)
}

// @Summary GetUserAccount
// @Tags user
// @Security ApiKeyAuth
// @Description get user account
// @Accept  json
// @Produce json
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} []dto.UserInfoDTO
// @Router /users/me [get]
func (h *Handler) findUserAccount(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	userInfo, err := h.userService.FindUserInfo(context.Request.Context(), userID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	userInfoDTO := dto.UserInfoDTO{
		Name:    userInfo.Name,
		Surname: userInfo.Surname,
	}

	h.successResponse(context, userInfoDTO)
}

// @Summary UpdateUser
// @Tags user
// @Security ApiKeyAuth
// @Description update user information
// @Accept  json
// @Produce json
// @Param input body dto.UpdateUserDTO true "update user info"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} []dto.UserDTO
// @Router /users/me [patch]
func (h *Handler) updateUser(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	var updateUserDTO dto.UpdateUserDTO
	err = context.ShouldBindJSON(&updateUserDTO)
	if err != nil {
		h.errorResponse(context, err)
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
	h.successResponse(context, userDTO)
}

// @Summary AddUserFreeCourse
// @Tags user
// @Security ApiKeyAuth
// @Description add user free course
// @Accept  json
// @Produce json
// @Param id path string true "course id"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {string} string
// @Router /users/me/courses/{id} [put]
func (h *Handler) addUserFreeCourse(context *gin.Context) {
	courseID, err := getIdFromPath(context, "course_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	course, err := h.courseService.FindByID(context.Request.Context(), courseID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	if course.Price > 0 {
		h.errorResponse(context, ForbiddenError)
		return
	}

	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	err = h.courseService.AddCourseStudent(context.Request.Context(), userID, courseID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, "successfully added free course")
}

// @Summary FindUserCourses
// @Tags user
// @Security ApiKeyAuth
// @Description find user courses
// @Accept  json
// @Produce json
// @Failure 400 {object} RestErrorBadRequest
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} []dto.CourseDTO
// @Router /users/me/courses [get]
func (h *Handler) findUserCourses(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	courses, err := h.courseService.FindStudentCourses(context.Request.Context(), userID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	courseDTOs := make([]dto.CourseDTO, len(courses))
	for i, course := range courses {
		courseDTOs[i] = dto.NewCourseDTO(course)
	}

	h.successResponse(context, courseDTOs)
}
