package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

func (h *Handler) initCourseRoutes(api *gin.RouterGroup) {
	courses := api.Group("/courses")
	{
		courses.GET("/", h.findAllCourses)
		courses.GET("/:id", h.findCourseByID)
		authenticated := courses.Group("/", h.verifyToken)
		{
			authenticated.GET("/:id/lessons", h.verifyCourseReadAccess, h.findCourseLessons)
			authenticated.POST("/:id/lessons", h.verifyCourseWriteAccess, h.createCourseLesson)
			authenticated.PUT("/:id/lessons/:lesson_id", h.verifyCourseWriteAccess, h.updateCourseLesson)
			authenticated.DELETE("/:id/lessons/:lesson_id", h.verifyCourseWriteAccess, h.deleteCourseLesson)

			authenticated.GET("/:id/teachers", h.findCourseTeachers)
			authenticated.PUT("/:id/teachers/:teacher_id", h.addCourseTeacher)

			authenticated.GET("/:id/reviews", h.findCourseReviews)
			authenticated.POST("/:id/reviews", h.addCourseReview)
		}
	}
}

func (h *Handler) findAllCourses(context *gin.Context) {
	courses, err := h.courseService.FindAll(context.Request.Context())
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

func (h *Handler) findCourseByID(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	course, err := h.courseService.FindByID(context.Request.Context(), courseID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	courseDTO := dto.NewCourseDTO(course)
	SuccessResponse(context, courseDTO)
}

func (h *Handler) findCourseTeachers(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	teachers, err := h.courseService.FindCourseTeachers(context.Request.Context(), courseID)
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

func (h *Handler) addCourseTeacher(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	course, err := h.courseService.FindByID(context.Request.Context(), courseID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	if !h.checkCurrentUserIsSchoolOwner(context, course.SchoolID) {
		ErrorResponse(context, ForbiddenError)
		return
	}

	teacherID, err := getIdFromPath(context, "teacher_id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	err = h.courseService.AddCourseTeacher(context.Request.Context(), teacherID, courseID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	SuccessResponse(context, "successfully added teacher")
}

func (h *Handler) findCourseLessons(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	lessons, err := h.lessonService.FindCourseLessons(context.Request.Context(), courseID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	lessonDTOs := make([]dto.LessonDTO, len(lessons))
	for i, lesson := range lessons {
		lessonDTOs[i] = dto.NewLessonDTO(lesson)
	}

	SuccessResponse(context, lessonDTOs)
}

func (h *Handler) createCourseLesson(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	var createLessonDTO dto.CreateLessonDTO
	err = context.ShouldBindJSON(&createLessonDTO)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	var lesson domain.Lesson
	switch createLessonDTO.Type {
	case dto.LessonDTOTheory:
		if !createLessonDTO.Theory.Valid {
			ErrorResponse(context, BadRequestError)
			return
		}
		lesson, err = h.lessonService.CreateTheoryLesson(context.Request.Context(),
			courseID, port.CreateTheoryParam{
				Title:  createLessonDTO.Title,
				Score:  int(createLessonDTO.Score.Int64),
				Theory: createLessonDTO.Theory.String,
			})
	case dto.LessonDTOVideo:
		if !createLessonDTO.VideoUrl.Valid {
			ErrorResponse(context, BadRequestError)
			return
		}
		lesson, err = h.lessonService.CreateVideoLesson(context.Request.Context(),
			courseID, port.CreateVideoParam{
				Title:    createLessonDTO.Title,
				Score:    int(createLessonDTO.Score.Int64),
				VideoUrl: createLessonDTO.VideoUrl.String,
			})
	case dto.LessonDTOPractice:
		if createLessonDTO.Tests == nil {
			ErrorResponse(context, BadRequestError)
			return
		}

		tests := make([]port.CreateTestParam, len(createLessonDTO.Tests))
		for i, test := range createLessonDTO.Tests {
			tests[i] = port.CreateTestParam{
				Task:    test.Task,
				Options: test.Options,
				Answer:  test.Answer,
				Level:   int(test.Level.Int64),
				Score:   int(test.Score.Int64),
			}
		}
		lesson, err = h.lessonService.CreatePracticeLesson(context.Request.Context(),
			courseID, port.CreatePracticeParam{
				Title: createLessonDTO.Title,
				Score: int(createLessonDTO.Score.Int64),
				Tests: tests,
			})
	default:
		ErrorResponse(context, BadRequestError)
		return
	}
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	lessonDTO := dto.NewLessonDTO(lesson)
	CreatedResponse(context, lessonDTO)
}

func (h *Handler) updateCourseLesson(context *gin.Context) {
	lessonID, err := getIdFromPath(context, "lesson_id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	var updateLessonDTO dto.UpdateLessonDTO
	err = context.ShouldBindJSON(&updateLessonDTO)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	lesson, err := h.lessonService.FindByID(context.Request.Context(), lessonID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	var updatedLesson domain.Lesson
	switch lesson.Type {
	case domain.TheoryLesson:
		if !updateLessonDTO.Theory.Valid {
			ErrorResponse(context, BadRequestError)
			return
		}
		updatedLesson, err = h.lessonService.UpdateTheoryLesson(context.Request.Context(),
			lessonID, port.UpdateTheoryParam{
				Title:  updateLessonDTO.Title,
				Score:  updateLessonDTO.Score,
				Theory: updateLessonDTO.Theory,
			})
	case domain.VideoLesson:
		if !updateLessonDTO.VideoUrl.Valid {
			ErrorResponse(context, BadRequestError)
			return
		}
		updatedLesson, err = h.lessonService.UpdateVideoLesson(context.Request.Context(),
			lessonID, port.UpdateVideoParam{
				Title:    updateLessonDTO.Title,
				Score:    updateLessonDTO.Score,
				VideoUrl: updateLessonDTO.VideoUrl,
			})
	case domain.PracticeLesson:
		if updateLessonDTO.Tests == nil {
			ErrorResponse(context, BadRequestError)
			return
		}

		tests := make([]port.UpdateTestParam, len(updateLessonDTO.Tests))
		for i, test := range updateLessonDTO.Tests {
			tests[i] = port.UpdateTestParam{
				Task:    test.Task,
				Options: test.Options,
				Answer:  test.Answer,
				Level:   int(test.Level.Int64),
				Score:   int(test.Score.Int64),
			}
		}
		updatedLesson, err = h.lessonService.UpdatePracticeLesson(context.Request.Context(),
			lessonID, port.UpdatePracticeParam{
				Title: updateLessonDTO.Title,
				Score: updateLessonDTO.Score,
				Tests: tests,
			})
	default:
		ErrorResponse(context, BadRequestError)
		return
	}
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	lessonDTO := dto.NewLessonDTO(updatedLesson)
	CreatedResponse(context, lessonDTO)
}

func (h *Handler) deleteCourseLesson(context *gin.Context) {
	lessonID, err := getIdFromPath(context, "lesson_id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	err = h.lessonService.Delete(context.Request.Context(), lessonID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	SuccessResponse(context, "lesson successfully deleted")
}

func (h *Handler) addCourseReview(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		ErrorResponse(context, UnauthorizedError)
		return
	}

	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	var createReviewDTO dto.CreateReviewDTO
	err = context.ShouldBindJSON(&createReviewDTO)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	review, err := h.reviewService.CreateCourseReview(context, courseID, userID,
		port.CreateReviewParam{Text: createReviewDTO.Text})

	reviewDTO := dto.NewReviewDTO(review)
	SuccessResponse(context, reviewDTO)
}

func (h *Handler) findCourseReviews(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	reviews, err := h.reviewService.FindCourseReviews(context.Request.Context(), courseID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	reviewDTOs := make([]dto.ReviewDTO, len(reviews))
	for i, review := range reviews {
		reviewDTOs[i] = dto.NewReviewDTO(review)
	}

	SuccessResponse(context, reviewDTOs)
}

func (h *Handler) verifyCourseWriteAccess(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	if !h.checkCurrentUserIsCourseTeacher(context, courseID) {
		ErrorResponse(context, ForbiddenError)
		return
	}
}

func (h *Handler) verifyCourseReadAccess(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	if !h.checkCurrentUserIsCourseOwner(context, courseID) &&
		!h.checkCurrentUserIsCourseTeacher(context, courseID) &&
		!h.checkCurrentUserIsCourseStudent(context, courseID) {
		ErrorResponse(context, ForbiddenError)
		return
	}
}

func (h *Handler) checkCurrentUserIsCourseStudent(context *gin.Context, courseID domain.ID) bool {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		ErrorResponse(context, UnauthorizedError)
		return false
	}

	isStudent, err := h.courseService.IsCourseStudent(context.Request.Context(), userID, courseID)
	if err != nil {
		ErrorResponse(context, err)
		return false
	}

	return isStudent
}

func (h *Handler) checkCurrentUserIsCourseTeacher(context *gin.Context, courseID domain.ID) bool {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		ErrorResponse(context, UnauthorizedError)
		return false
	}

	isTeacher, err := h.courseService.IsCourseTeacher(context.Request.Context(), userID, courseID)
	if err != nil {
		ErrorResponse(context, err)
		return false
	}

	return isTeacher
}

func (h *Handler) checkCurrentUserIsCourseOwner(context *gin.Context, courseID domain.ID) bool {
	course, err := h.courseService.FindByID(context.Request.Context(), courseID)
	if err != nil {
		ErrorResponse(context, err)
		return false
	}

	return h.checkCurrentUserIsSchoolOwner(context, course.SchoolID)
}
