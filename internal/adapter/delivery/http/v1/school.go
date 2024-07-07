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
			authenticated.PUT("/:id", h.verifySchoolOwner, h.updateSchool)

			authenticated.GET("/:id/courses", h.findSchoolCourses)
			authenticated.POST("/:id/courses", h.verifySchoolOwner, h.createSchoolCourse)
			authenticated.PUT("/:id/courses/:course_id", h.verifySchoolOwner, h.updateSchoolCourse)
			authenticated.DELETE("/:id/courses/:course_id", h.verifySchoolOwner, h.deleteSchoolCourse)

			authenticated.GET("/:id/teachers", h.findSchoolTeachers)
			// https://datatracker.ietf.org/doc/html/rfc2616#section-9.6
			// If the Request-URI does not point to an existing resource,
			// and that URI is capable of being defined as a new resource by
			// the requesting user agent, the origin server can create the resource with that URI.
			authenticated.PUT("/:id/teachers/:teacher_id", h.verifySchoolOwner, h.addSchoolTeacher)
		}
	}
}

func (h *Handler) findAllSchools(context *gin.Context) {
	schools, err := h.schoolService.FindAll(context.Request.Context())
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	schoolDTOs := make([]dto.SchoolDTO, len(schools))
	for i, school := range schools {
		schoolDTOs[i] = dto.NewSchoolDTO(school)
	}

	h.successResponse(context, schoolDTOs)
}

func (h *Handler) findSchoolByID(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	school, err := h.schoolService.FindByID(context.Request.Context(), schoolID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	schoolDTO := dto.NewSchoolDTO(school)
	h.successResponse(context, schoolDTO)
}

func (h *Handler) createSchool(context *gin.Context) {
	var createSchoolDTO dto.CreateSchoolDTO
	err := context.ShouldBindJSON(&createSchoolDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	school, err := h.schoolService.CreateUserSchool(context.Request.Context(), userID,
		port.CreateSchoolParam{
			Name:        createSchoolDTO.Name,
			Description: createSchoolDTO.Description.String,
		})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	schoolDTO := dto.NewSchoolDTO(school)
	h.createdResponse(context, schoolDTO)
}

func (h *Handler) updateSchool(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var updateSchoolDTO dto.UpdateSchoolDTO
	err = context.ShouldBindJSON(&updateSchoolDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	school, err := h.schoolService.Update(context.Request.Context(),
		schoolID, port.UpdateSchoolParam{
			Description: updateSchoolDTO.Description,
		})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	schoolDTO := dto.NewSchoolDTO(school)
	h.successResponse(context, schoolDTO)
}

func (h *Handler) findSchoolCourses(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	courses, err := h.schoolService.FindSchoolCourses(context.Request.Context(), schoolID)
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

func (h *Handler) createSchoolCourse(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var createCourseDTO dto.CreateCourseDTO
	err = context.ShouldBindJSON(&createCourseDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	course, err := h.courseService.CreateSchoolCourse(context.Request.Context(),
		schoolID, port.CreateCourseParam{
			Name:     createCourseDTO.Name,
			Level:    int(createCourseDTO.Level.Int64),
			Price:    createCourseDTO.Price.Int64,
			Language: createCourseDTO.Language,
		})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	courseDTO := dto.NewCourseDTO(course)
	h.createdResponse(context, courseDTO)
}

func (h *Handler) updateSchoolCourse(context *gin.Context) {
	courseID, err := getIdFromPath(context, "course_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var updateCourseDTO dto.UpdateCourseDTO
	err = context.ShouldBindJSON(&updateCourseDTO)
	if err != nil {
		h.errorResponse(context, err)
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
		h.errorResponse(context, err)
		return
	}

	courseDTO := dto.NewCourseDTO(course)
	h.successResponse(context, courseDTO)
}

func (h *Handler) deleteSchoolCourse(context *gin.Context) {
	courseID, err := getIdFromPath(context, "course_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.courseService.Delete(context.Request.Context(), courseID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, "successfully deleted")
}

func (h *Handler) findSchoolTeachers(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	teachers, err := h.schoolService.FindSchoolTeachers(context.Request.Context(), schoolID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	teacherDTOs := make([]dto.UserDTO, len(teachers))
	for i, teacher := range teachers {
		teacherDTOs[i] = dto.NewUserDTO(teacher)
	}

	h.successResponse(context, teacherDTOs)
}

func (h *Handler) addSchoolTeacher(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	teacherID, err := getIdFromPath(context, "teacher_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.schoolService.AddSchoolTeacher(context.Request.Context(), schoolID, teacherID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.createdResponse(context, "teacher successfully added")
}

func (h *Handler) verifySchoolOwner(context *gin.Context) {
	schoolID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	if !h.checkCurrentUserIsSchoolOwner(context, schoolID) {
		h.errorResponse(context, ForbiddenError)
		return
	}
}

func (h *Handler) checkCurrentUserIsSchoolOwner(context *gin.Context, schoolID domain.ID) bool {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		return false
	}

	school, err := h.schoolService.FindByID(context.Request.Context(), schoolID)
	if err != nil {
		h.errorResponse(context, err)
		return false
	}

	return school.OwnerID == userID
}
