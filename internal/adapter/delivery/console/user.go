package console

import (
	"context"
	"fmt"
	"github.com/paw1a/eschool/internal/adapter/delivery/console/dto"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

//func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
//	users := api.Group("/users")
//	{
//		users.GET("/", h.findAllUsers)
//		authenticated := users.Group("/", h.verifyToken)
//		{
//			authenticated.GET("/account", h.findUserAccount)
//			authenticated.PUT("/account", h.updateUser)
//
//			authenticated.GET("/courses", h.findUserCourses)
//			authenticated.PUT("/courses/:course_id", h.addUserFreeCourse)
//			authenticated.GET("/certificates", h.findUserCertificates)
//		}
//	}
//}

func (h *Handler) GetAllUsers(c *Console) {
	users, err := h.userService.FindAll(context.Background())
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(users) == 0 {
		fmt.Println("no users")
		return
	}

	for _, user := range users {
		dto.PrintUserDTO(dto.NewUserDTO(user))
		fmt.Println()
	}
}

func (h *Handler) GetUserAccount(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}
	userID := *c.UserID

	user, err := h.userService.FindByID(context.Background(), userID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Printf("ID: %s\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Surname: %s\n", user.Surname)
}

func (h *Handler) UpdateUser(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}
	userID := *c.UserID

	var updateUserDTO dto.UpdateUserDTO
	err = dto.InputUpdateUserDTO(&updateUserDTO)
	if err != nil {
		ErrorResponse(err)
		return
	}

	user, err := h.userService.Update(context.Background(), userID, port.UpdateUserParam{
		Name:      updateUserDTO.Name,
		Surname:   updateUserDTO.Surname,
		Phone:     updateUserDTO.Phone,
		City:      updateUserDTO.City,
		AvatarUrl: updateUserDTO.AvatarUrl,
	})

	userDTO := dto.NewUserDTO(user)
	dto.PrintUserDTO(userDTO)
}

func (h *Handler) AddUserFreeCourse(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}
	userID := *c.UserID

	var courseID domain.ID
	err = dto.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	course, err := h.courseService.FindByID(context.Background(), courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if course.Price > 0 {
		ErrorResponse(ForbiddenError)
		return
	}

	err = h.courseService.AddCourseStudent(context.Background(), userID, courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("successfully added free course")
}

func (h *Handler) FindUserCourses(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}
	userID := *c.UserID

	courses, err := h.courseService.FindStudentCourses(context.Background(), userID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(courses) == 0 {
		fmt.Println("no courses")
		return
	}

	for _, course := range courses {
		dto.PrintCourseDTO(dto.NewCourseDTO(course))
		fmt.Println()
	}
}

//func (h *Handler) FindUserCertificates(c *Console) {
//	userID, err := getIdFromRequestContext(context)
//	if err != nil {
//		ErrorResponse(context, UnauthorizedError)
//		return
//	}
//
//	courses, err := h.courseService.FindStudentCourses(context.Request.Context(), userID)
//	if err != nil {
//		ErrorResponse(context, err)
//		return
//	}
//
//	courseDTOs := make([]dto.CourseDTO, len(courses))
//	for i, course := range courses {
//		courseDTOs[i] = dto.NewCourseDTO(course)
//	}
//
//	SuccessResponse(context, courseDTOs)
//}
