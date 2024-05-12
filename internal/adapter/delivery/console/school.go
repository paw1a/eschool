package console

import (
	"bufio"
	"context"
	"fmt"
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/adapter/delivery/console/dto"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"os"
	"strings"
)

func (h *Handler) FindAllSchools(c *Console) {
	schools, err := h.schoolService.FindAll(context.Background())
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(schools) == 0 {
		fmt.Println("no schools")
		return
	}

	for _, school := range schools {
		dto.PrintSchoolDTO(dto.NewSchoolDTO(school))
		fmt.Println()
	}
}

func (h *Handler) FindSchoolByID(c *Console) {
	var schoolID domain.ID
	err := dto.InputID(&schoolID, "school")
	if err != nil {
		ErrorResponse(err)
		return
	}

	school, err := h.schoolService.FindByID(context.Background(), schoolID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	schoolDTO := dto.NewSchoolDTO(school)
	dto.PrintSchoolDTO(schoolDTO)
}

func (h *Handler) CreateSchool(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}
	userID := *c.UserID

	var createSchoolDTO dto.CreateSchoolDTO
	err = dto.InputCreateSchoolDTO(&createSchoolDTO)
	if err != nil {
		ErrorResponse(err)
		return
	}

	school, err := h.schoolService.CreateUserSchool(context.Background(), userID,
		port.CreateSchoolParam{
			Name:        createSchoolDTO.Name,
			Description: createSchoolDTO.Description.String,
		})
	if err != nil {
		ErrorResponse(err)
		return
	}

	schoolDTO := dto.NewSchoolDTO(school)
	dto.PrintSchoolDTO(schoolDTO)
}

func (h *Handler) UpdateSchool(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}

	var schoolID domain.ID
	err = dto.InputID(&schoolID, "school")
	if err != nil {
		ErrorResponse(err)
		return
	}

	var updateSchoolDTO dto.UpdateSchoolDTO
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("School description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	if description != "" {
		updateSchoolDTO.Description = null.StringFrom(description)
	}

	school, err := h.schoolService.Update(context.Background(),
		schoolID, port.UpdateSchoolParam{
			Description: updateSchoolDTO.Description,
		})
	if err != nil {
		ErrorResponse(err)
		return
	}

	schoolDTO := dto.NewSchoolDTO(school)
	dto.PrintSchoolDTO(schoolDTO)
}

func (h *Handler) FindSchoolCourses(c *Console) {
	var schoolID domain.ID
	err := dto.InputID(&schoolID, "school")
	if err != nil {
		ErrorResponse(err)
		return
	}

	courses, err := h.schoolService.FindSchoolCourses(context.Background(), schoolID)
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

func (h *Handler) CreateSchoolCourse(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}

	var schoolID domain.ID
	err = dto.InputID(&schoolID, "school")
	if err != nil {
		ErrorResponse(err)
		return
	}

	if !h.verifySchoolOwner(c, schoolID) {
		ErrorResponse(ForbiddenError)
		return
	}

	var createCourseDTO dto.CreateCourseDTO
	err = dto.InputCreateCourseDTO(&createCourseDTO)
	if err != nil {
		ErrorResponse(err)
		return
	}

	course, err := h.courseService.CreateSchoolCourse(context.Background(),
		schoolID, port.CreateCourseParam{
			Name:     createCourseDTO.Name,
			Level:    int(createCourseDTO.Level.Int64),
			Price:    createCourseDTO.Price.Int64,
			Language: createCourseDTO.Language,
		})
	if err != nil {
		ErrorResponse(err)
		return
	}

	courseDTO := dto.NewCourseDTO(course)
	dto.PrintCourseDTO(courseDTO)
}

func (h *Handler) UpdateSchoolCourse(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}

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

	if !h.verifySchoolOwner(c, course.SchoolID) {
		ErrorResponse(ForbiddenError)
		return
	}

	var updateCourseDTO dto.UpdateCourseDTO
	err = dto.InputUpdateCourseDTO(&updateCourseDTO)
	if err != nil {
		ErrorResponse(err)
		return
	}

	updatedCourse, err := h.courseService.Update(context.Background(),
		courseID, port.UpdateCourseParam{
			Name:     updateCourseDTO.Name,
			Level:    updateCourseDTO.Level,
			Price:    updateCourseDTO.Price,
			Language: updateCourseDTO.Language,
		})
	if err != nil {
		ErrorResponse(err)
		return
	}

	courseDTO := dto.NewCourseDTO(updatedCourse)
	dto.PrintCourseDTO(courseDTO)
}

func (h *Handler) DeleteSchoolCourse(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}

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

	if !h.verifySchoolOwner(c, course.SchoolID) {
		ErrorResponse(ForbiddenError)
		return
	}

	err = h.courseService.Delete(context.Background(), courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}
	fmt.Println("successfully deleted")
}

func (h *Handler) FindSchoolTeachers(c *Console) {
	var schoolID domain.ID
	err := dto.InputID(&schoolID, "school")
	if err != nil {
		ErrorResponse(err)
		return
	}

	teachers, err := h.schoolService.FindSchoolTeachers(context.Background(), schoolID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(teachers) == 0 {
		fmt.Println("no teachers")
		return
	}

	for _, teacher := range teachers {
		dto.PrintUserDTO(dto.NewUserDTO(teacher))
		fmt.Println()
	}
}

func (h *Handler) AddSchoolTeacher(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}

	var schoolID domain.ID
	err = dto.InputID(&schoolID, "school")
	if err != nil {
		ErrorResponse(err)
		return
	}

	if !h.verifySchoolOwner(c, schoolID) {
		ErrorResponse(ForbiddenError)
		return
	}

	var teacherID domain.ID
	err = dto.InputID(&teacherID, "teacher")
	if err != nil {
		ErrorResponse(err)
		return
	}

	err = h.schoolService.AddSchoolTeacher(context.Background(), schoolID, teacherID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("successfully added teacher")
}

func (h *Handler) verifySchoolOwner(c *Console, schoolID domain.ID) bool {
	return h.checkCurrentUserIsSchoolOwner(c, schoolID)
}

func (h *Handler) checkCurrentUserIsSchoolOwner(c *Console, schoolID domain.ID) bool {
	err := h.verifyAuth(c)
	if err != nil {
		return false
	}
	userID := *c.UserID

	school, err := h.schoolService.FindByID(context.Background(), schoolID)
	if err != nil {
		ErrorResponse(err)
		return false
	}

	return school.OwnerID.String() == userID.String()
}
