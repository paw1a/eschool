package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
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
			authenticated.GET("/:id/lessons/:lesson_id", h.verifyCourseReadAccess, h.findLessonByID)
			authenticated.POST("/:id/lessons", h.verifyCourseWriteAccess, h.createCourseLesson)
			authenticated.PUT("/:id/lessons/:lesson_id", h.verifyCourseWriteAccess, h.updateCourseLesson)
			authenticated.DELETE("/:id/lessons/:lesson_id", h.verifyCourseWriteAccess, h.deleteCourseLesson)

			authenticated.GET("/:id/teachers", h.findCourseTeachers)
			authenticated.PUT("/:id/teachers/:teacher_id", h.addCourseTeacher)

			authenticated.GET("/:id/reviews", h.findCourseReviews)
			authenticated.POST("/:id/reviews", h.addCourseReview)

			authenticated.GET("/:id/lessons/:lesson_id/stat", h.verifyCourseReadAccess, h.findLessonStat)
			authenticated.POST("/:id/lessons/:lesson_id/stat", h.verifyCourseReadAccess, h.passCourseLesson)
		}
	}
}

// @Summary GetAllCourses
// @Tags course
// @Description get all courses
// @Accept  json
// @Produce json
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} []dto.CourseDTO
// @Router /courses [get]
func (h *Handler) findAllCourses(context *gin.Context) {
	courses, err := h.courseService.FindAll(context.Request.Context())
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

// @Summary GetCourseByID
// @Tags course
// @Description get course by id
// @Accept  json
// @Produce json
// @Param   id   path    string  true  "course id"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} []dto.CourseDTO
// @Router /courses/{id} [get]
func (h *Handler) findCourseByID(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	course, err := h.courseService.FindByID(context.Request.Context(), courseID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	courseDTO := dto.NewCourseDTO(course)
	h.successResponse(context, courseDTO)
}

// @Summary GetLessonByID
// @Tags course
// @Security ApiKeyAuth
// @Description get lesson by id
// @Accept  json
// @Produce json
// @Param   courseID   path    string  true  "course id"
// @Param   lessonID   path    string  true  "lesson id"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} []dto.LessonDTO
// @Router /courses/{courseID}/lessons/{lessonID} [get]
func (h *Handler) findLessonByID(context *gin.Context) {
	lessonID, err := getIdFromPath(context, "lesson_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	lesson, err := h.lessonService.FindByID(context.Request.Context(), lessonID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	lessonDTO := dto.NewLessonDTO(lesson)
	h.successResponse(context, lessonDTO)
}

// @Summary GetCourseTeachers
// @Tags course
// @Description get course teachers
// @Accept  json
// @Produce json
// @Param   id   path    string  true  "course id"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} []dto.UserDTO
// @Router /courses/{id}/teachers [get]
func (h *Handler) findCourseTeachers(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	teachers, err := h.courseService.FindCourseTeachers(context.Request.Context(), courseID)
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

// @Summary AddCourseTeacher
// @Tags course
// @Security ApiKeyAuth
// @Description add course teacher
// @Accept  json
// @Produce json
// @Param   courseID    path    string  true  "course id"
// @Param   teacherID   path    string  true  "course id"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 201 {string} string "message"
// @Router /courses/{courseID}/teachers/{teacherID} [put]
func (h *Handler) addCourseTeacher(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	course, err := h.courseService.FindByID(context.Request.Context(), courseID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	if !h.checkCurrentUserIsSchoolOwner(context, course.SchoolID) {
		h.errorResponse(context, ForbiddenError)
		return
	}

	teacherID, err := getIdFromPath(context, "teacher_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.courseService.AddCourseTeacher(context.Request.Context(), teacherID, courseID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, "successfully added teacher")
}

// @Summary GetCourseLessons
// @Tags course
// @Security ApiKeyAuth
// @Description get course lessons
// @Accept  json
// @Produce json
// @Param   id   path    string  true  "course id"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} []dto.LessonDTO
// @Router /courses/{id}/lessons [get]
func (h *Handler) findCourseLessons(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	lessons, err := h.lessonService.FindCourseLessons(context.Request.Context(), courseID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	lessonDTOs := make([]dto.LessonDTO, len(lessons))
	for i, lesson := range lessons {
		lessonDTOs[i] = dto.NewLessonDTO(lesson)
	}

	h.successResponse(context, lessonDTOs)
}

// @Summary CreateCourseLesson
// @Tags course
// @Security ApiKeyAuth
// @Description create course lesson
// @Accept  json
// @Produce json
// @Param   id   path    string  true  "course id"
// @Param input body dto.CreateLessonDTO true "created lesson info"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 201 {object} dto.LessonDTO
// @Router /courses/{id}/lessons [post]
func (h *Handler) createCourseLesson(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var createLessonDTO dto.CreateLessonDTO
	err = context.ShouldBindJSON(&createLessonDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var lesson domain.Lesson
	switch createLessonDTO.Type {
	case dto.LessonDTOTheory:
		if !createLessonDTO.Theory.Valid {
			h.errorResponse(context, BadRequestError)
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
			h.errorResponse(context, BadRequestError)
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
			h.errorResponse(context, BadRequestError)
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
		h.errorResponse(context, BadRequestError)
		return
	}
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	lessonDTO := dto.NewLessonDTO(lesson)
	h.createdResponse(context, lessonDTO)
}

// @Summary UpdateCourseLesson
// @Tags course
// @Security ApiKeyAuth
// @Description update course lesson
// @Accept  json
// @Produce json
// @Param courseID   path    string  true  "course id"
// @Param lessonID   path    string  true  "lesson id"
// @Param input body dto.UpdateLessonDTO true "updated lesson info"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} dto.LessonDTO
// @Router /courses/{courseID}/lessons/{lessonID} [put]
func (h *Handler) updateCourseLesson(context *gin.Context) {
	lessonID, err := getIdFromPath(context, "lesson_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var updateLessonDTO dto.UpdateLessonDTO
	err = context.ShouldBindJSON(&updateLessonDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	lesson, err := h.lessonService.FindByID(context.Request.Context(), lessonID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var updatedLesson domain.Lesson
	switch lesson.Type {
	case domain.TheoryLesson:
		if !updateLessonDTO.Theory.Valid {
			h.errorResponse(context, BadRequestError)
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
			h.errorResponse(context, BadRequestError)
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
			h.errorResponse(context, BadRequestError)
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
		h.errorResponse(context, BadRequestError)
		return
	}
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	lessonDTO := dto.NewLessonDTO(updatedLesson)
	h.createdResponse(context, lessonDTO)
}

// @Summary DeleteCourseLesson
// @Tags course
// @Security ApiKeyAuth
// @Description delete course lesson
// @Accept  json
// @Produce json
// @Param   courseID   path    string  true  "course id"
// @Param   lessonID   path    string  true  "lesson id"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {string} string "message"
// @Router /schools/{courseID}/courses/{lessonID} [delete]
func (h *Handler) deleteCourseLesson(context *gin.Context) {
	lessonID, err := getIdFromPath(context, "lesson_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.lessonService.Delete(context.Request.Context(), lessonID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, "lesson successfully deleted")
}

func (h *Handler) addCourseReview(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var createReviewDTO dto.CreateReviewDTO
	err = context.ShouldBindJSON(&createReviewDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	review, err := h.reviewService.CreateCourseReview(context, courseID, userID,
		port.CreateReviewParam{Text: createReviewDTO.Text})

	reviewDTO := dto.NewReviewDTO(review)
	h.successResponse(context, reviewDTO)
}

func (h *Handler) findCourseReviews(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	reviews, err := h.reviewService.FindCourseReviews(context.Request.Context(), courseID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	reviewDTOs := make([]dto.ReviewDTO, len(reviews))
	for i, review := range reviews {
		reviewDTOs[i] = dto.NewReviewDTO(review)
	}

	h.successResponse(context, reviewDTOs)
}

// @Summary GetLessonStat
// @Tags course
// @Security ApiKeyAuth
// @Description get lesson statistics
// @Accept  json
// @Produce json
// @Param   courseID   path    string  true  "course id"
// @Param   lessonID   path    string  true  "lesson id"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} dto.LessonStatDTO
// @Router /schools/{courseID}/courses/{lessonID}/stat [get]
func (h *Handler) findLessonStat(context *gin.Context) {
	lessonID, err := getIdFromPath(context, "lesson_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	stat, err := h.statService.FindLessonStat(context.Request.Context(), userID, lessonID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	statDTO := dto.NewLessonStatDTO(stat)
	h.successResponse(context, statDTO)
}

// @Summary PassCourseLesson
// @Tags course
// @Security ApiKeyAuth
// @Description pass course lesson
// @Accept  json
// @Produce json
// @Param   courseID   path    string  true  "course id"
// @Param   lessonID   path    string  true  "lesson id"
// @Param input body dto.PassLessonDTO true "passed lesson info"
// @Failure 400 {object} RestErrorBadRequest
// @Failure 401 {object} RestErrorUnauthorized
// @Failure 404 {object} RestErrorNotFound
// @Failure 500 {object} RestErrorInternalError
// @Success 200 {object} dto.LessonStatDTO
// @Router /schools/{courseID}/courses/{lessonID}/stat [post]
func (h *Handler) passCourseLesson(context *gin.Context) {
	lessonID, err := getIdFromPath(context, "lesson_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	var passLessonDTO dto.PassLessonDTO
	err = context.ShouldBindJSON(&passLessonDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	lesson, err := h.lessonService.FindByID(context, lessonID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	switch lesson.Type {
	case domain.TheoryLesson:
		fallthrough
	case domain.VideoLesson:
		err = h.statService.UpdateLessonStat(context.Request.Context(), userID,
			lessonID, port.UpdateLessonStatParam{
				Score:     null.IntFrom(int64(lesson.Score)),
				TestStats: nil,
			})
	case domain.PracticeLesson:
		if len(passLessonDTO.PassTests) == 0 {
			h.errorResponse(context, BadRequestError)
			return
		}

		testStats := make([]port.UpdateTestStatParam, len(passLessonDTO.PassTests))
		for i, passTest := range passLessonDTO.PassTests {
			for j, test := range lesson.Tests {
				if passTest.TestID == test.ID.String() {
					var newScore int
					if passTest.Answer == test.Answer {
						newScore = lesson.Tests[j].Score
					}

					testStats[i] = port.UpdateTestStatParam{
						TestID: test.ID,
						Score:  newScore,
					}
				}
			}
		}

		err = h.statService.UpdateLessonStat(context.Request.Context(), userID,
			lessonID, port.UpdateLessonStatParam{
				Score:     null.IntFrom(int64(lesson.Score)),
				TestStats: testStats,
			})
	}

	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, "successfully passed lesson")
}

func (h *Handler) verifyCourseWriteAccess(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	if !h.checkCurrentUserIsCourseTeacher(context, courseID) {
		h.errorResponse(context, ForbiddenError)
		return
	}
}

func (h *Handler) verifyCourseReadAccess(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	if !h.checkCurrentUserIsCourseOwner(context, courseID) &&
		!h.checkCurrentUserIsCourseTeacher(context, courseID) &&
		!h.checkCurrentUserIsCourseStudent(context, courseID) {
		h.errorResponse(context, ForbiddenError)
		return
	}
}

func (h *Handler) checkCurrentUserIsCourseStudent(context *gin.Context, courseID domain.ID) bool {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		return false
	}

	isStudent, err := h.courseService.IsCourseStudent(context.Request.Context(), userID, courseID)
	if err != nil {
		h.errorResponse(context, err)
		return false
	}

	return isStudent
}

func (h *Handler) checkCurrentUserIsCourseTeacher(context *gin.Context, courseID domain.ID) bool {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		return false
	}

	isTeacher, err := h.courseService.IsCourseTeacher(context.Request.Context(), userID, courseID)
	if err != nil {
		h.errorResponse(context, err)
		return false
	}

	return isTeacher
}

func (h *Handler) checkCurrentUserIsCourseOwner(context *gin.Context, courseID domain.ID) bool {
	course, err := h.courseService.FindByID(context.Request.Context(), courseID)
	if err != nil {
		h.errorResponse(context, err)
		return false
	}

	return h.checkCurrentUserIsSchoolOwner(context, course.SchoolID)
}
