package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

func (h *Handler) initSchoolRoutes(api *gin.RouterGroup) {
	schools := api.Group("/schools")
	{
		schools.GET("/", h.findAllSchools)
		schools.GET("/:id", h.findSchoolByID)
		authenticated := schools.Group("/", h.verifyToken)
		{
			authenticated.POST("/", h.createSchool)
			authenticated.PUT("/:id", h.updateSchool)

			authenticated.GET("/:id/courses", h.findSchoolCourses)
			authenticated.POST("/:id/courses", h.createSchoolCourse)
			authenticated.PUT("/:school_id/courses/:course_id", h.updateSchoolCourse)
			authenticated.DELETE("/:school_id/courses/:course_id", h.deleteSchoolCourse)

			authenticated.GET("/:id/teachers", h.findSchoolTeachers)
			authenticated.POST("/:id/teachers", h.addSchoolTeacher)
		}
	}
}

func (h *Handler) findAllSchools(context *gin.Context) {
	schools, err := h.schoolService.FindAll(context.Request.Context())
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	schoolDTOs := make([]dto.SchoolDTO, len(schools))
	for i, school := range schools {
		schoolDTOs[i] = dto.NewSchoolDTO(school)
	}

	SuccessResponse(context, schoolDTOs)
}

func (h *Handler) findSchoolByID(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	school, err := h.schoolService.FindByID(context.Request.Context(), schoolID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	schoolDTO := dto.NewSchoolDTO(school)
	SuccessResponse(context, schoolDTO)
}

func (h *Handler) createSchool(context *gin.Context) {
	var createSchoolDTO dto.CreateSchoolDTO
	err := context.BindJSON(&createSchoolDTO)
	if err != nil {
		ErrorResponse(context, UnmarshalError)
		return
	}

	userID, err := getIdFromRequestContext(context)
	if err != nil {
		ErrorResponse(context, UnauthorizedError)
		return
	}

	school, err := h.schoolService.CreateUserSchool(context.Request.Context(), userID,
		port.CreateSchoolParam{
			Name:        createSchoolDTO.Name,
			Description: createSchoolDTO.Description,
		})
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	schoolDTO := dto.NewSchoolDTO(school)
	SuccessResponse(context, schoolDTO)
}

func (h *Handler) updateSchool(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	if !h.checkCurrentUserIsSchoolOwner(context, schoolID) {
		return
	}

	var updateSchoolDTO dto.UpdateSchoolDTO
	err = context.BindJSON(&updateSchoolDTO)
	if err != nil {
		ErrorResponse(context, UnmarshalError)
		return
	}

	school, err := h.schoolService.Update(context.Request.Context(),
		schoolID, port.UpdateSchoolParam{
			Description: updateSchoolDTO.Description,
		})
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	schoolDTO := dto.NewSchoolDTO(school)
	SuccessResponse(context, schoolDTO)
}

func (h *Handler) findSchoolCourses(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	courses, err := h.schoolService.FindSchoolCourses(context.Request.Context(), schoolID)
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

func (h *Handler) checkCurrentUserIsSchoolOwner(context *gin.Context, schoolID domain.ID) bool {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		ErrorResponse(context, UnauthorizedError)
		return false
	}

	school, err := h.schoolService.FindByID(context.Request.Context(), schoolID)
	if err != nil {
		ErrorResponse(context, err)
		return false
	}

	if school.OwnerID != userID {
		ErrorResponse(context, ForbiddenError)
		return false
	}

	return true
}

func (h *Handler) createSchoolCourse(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	if !h.checkCurrentUserIsSchoolOwner(context, schoolID) {
		return
	}

	var createCourseDTO dto.CreateCourseDTO
	err = context.BindJSON(&createCourseDTO)
	if err != nil {
		ErrorResponse(context, UnmarshalError)
		return
	}

	course, err := h.courseService.CreateSchoolCourse(context.Request.Context(),
		schoolID, port.CreateCourseParam{
			Name:     createCourseDTO.Name,
			Level:    createCourseDTO.Level,
			Price:    createCourseDTO.Price,
			Language: createCourseDTO.Language,
		})
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	courseDTO := dto.NewCourseDTO(course)
	SuccessResponse(context, courseDTO)
}

func (h *Handler) updateSchoolCourse(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "school_id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	courseID, err := getIdFromPath(context, "course_id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	if !h.checkCurrentUserIsSchoolOwner(context, schoolID) {
		return
	}

	var updateCourseDTO dto.UpdateCourseDTO
	err = context.BindJSON(&updateCourseDTO)
	if err != nil {
		ErrorResponse(context, UnmarshalError)
		return
	}

	course, err := h.courseService.Update(context.Request.Context(),
		courseID, port.UpdateCourseParam{
			Name:     updateCourseDTO.Name,
			Level:    updateCourseDTO.Level,
			Price:    updateCourseDTO.Price,
			Language: updateCourseDTO.Language,
		})
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	courseDTO := dto.NewCourseDTO(course)
	SuccessResponse(context, courseDTO)
}

func (h *Handler) deleteSchoolCourse(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "school_id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	courseID, err := getIdFromPath(context, "course_id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	if !h.checkCurrentUserIsSchoolOwner(context, schoolID) {
		return
	}

	err = h.courseService.Delete(context.Request.Context(), courseID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	SuccessResponse(context, "successfully deleted")
}

func (h *Handler) findSchoolTeachers(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	teachers, err := h.schoolService.FindSchoolTeachers(context.Request.Context(), schoolID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	teacherDTOs := make([]dto.UserDTO, len(teachers))
	for i, teacher := range teachers {
		teacherDTOs[i] = dto.NewUserDTO(teacher)
	}

	SuccessResponse(context, teacherDTOs)
}

func (h *Handler) addSchoolTeacher(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	if !h.checkCurrentUserIsSchoolOwner(context, schoolID) {
		return
	}

	var addTeacherDTO dto.AddTeacherDTO
	err = context.BindJSON(&addTeacherDTO)
	if err != nil {
		ErrorResponse(context, UnmarshalError)
		return
	}

	teacherID := domain.ID(addTeacherDTO.TeacherID)
	err = h.schoolService.AddSchoolTeacher(context.Request.Context(), schoolID, teacherID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	SuccessResponse(context, "successfully added")
}
