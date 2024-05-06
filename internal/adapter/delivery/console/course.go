package console

import (
	"context"
	"fmt"
	"github.com/paw1a/eschool/internal/adapter/delivery/console/dto"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

//func (h *Handler) initCourseRoutes(api *gin.RouterGroup) {
//	courses := api.Group("/courses")
//	{
//		courses.GET("/", h.findAllCourses)
//		courses.GET("/:id", h.findCourseByID)
//		authenticated := courses.Group("/", h.verifyToken)
//		{
//			authenticated.GET("/:id/lessons", h.verifyCourseReadAccess, h.findCourseLessons)
//			authenticated.GET("/:id/lessons/:lesson_id", h.verifyCourseReadAccess, h.findLessonByID)
//			authenticated.POST("/:id/lessons", h.verifyCourseWriteAccess, h.createCourseLesson)
//			authenticated.PUT("/:id/lessons/:lesson_id", h.verifyCourseWriteAccess, h.updateCourseLesson)
//			authenticated.DELETE("/:id/lessons/:lesson_id", h.verifyCourseWriteAccess, h.deleteCourseLesson)
//
//			authenticated.GET("/:id/teachers", h.findCourseTeachers)
//			authenticated.PUT("/:id/teachers/:teacher_id", h.addCourseTeacher)
//
//			authenticated.GET("/:id/reviews", h.findCourseReviews)
//			authenticated.POST("/:id/reviews", h.addCourseReview)
//
//			authenticated.GET("/:id/lessons/:lesson_id/stat", h.verifyCourseReadAccess, h.findLessonStat)
//			authenticated.POST("/:id/lessons/:lesson_id/stat", h.verifyCourseReadAccess, h.passCourseLesson)
//		}
//	}
//}

func (h *Handler) FindAllCourses(c *Console) {
	courses, err := h.courseService.FindAll(context.Background())
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

func (h *Handler) FindCourseByID(c *Console) {
	var courseID domain.ID
	err := dto.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	course, err := h.courseService.FindByID(context.Background(), courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	courseDTO := dto.NewCourseDTO(course)
	dto.PrintCourseDTO(courseDTO)
}

func (h *Handler) FindLessonByID(c *Console) {
	var lessonID domain.ID
	err := dto.InputID(&lessonID, "lesson")
	if err != nil {
		ErrorResponse(err)
		return
	}

	lesson, err := h.lessonService.FindByID(context.Background(), lessonID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if !h.verifyCourseReadAccess(c, lesson.CourseID) {
		ErrorResponse(ForbiddenError)
		return
	}

	lessonDTO := dto.NewLessonDTO(lesson)
	dto.PrintLessonDTO(lessonDTO)
}

func (h *Handler) FindCourseTeachers(c *Console) {
	var courseID domain.ID
	err := dto.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	teachers, err := h.courseService.FindCourseTeachers(context.Background(), courseID)
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

func (h *Handler) AddCourseTeacher(c *Console) {
	var courseID domain.ID
	err := dto.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	course, err := h.courseService.FindByID(context.Background(), courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if !h.checkCurrentUserIsSchoolOwner(c, course.SchoolID) {
		ErrorResponse(ForbiddenError)
		return
	}

	var teacherID domain.ID
	err = dto.InputID(&teacherID, "teacher")
	if err != nil {
		ErrorResponse(err)
		return
	}

	err = h.courseService.AddCourseTeacher(context.Background(), teacherID, courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("successfully added teacher")
}

func (h *Handler) FindCourseLessons(c *Console) {
	var courseID domain.ID
	err := dto.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	if !h.verifyCourseReadAccess(c, courseID) {
		ErrorResponse(ForbiddenError)
		return
	}

	lessons, err := h.lessonService.FindCourseLessons(context.Background(), courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(lessons) == 0 {
		fmt.Println("no lessons")
		return
	}

	for _, lesson := range lessons {
		dto.PrintLessonDTO(dto.NewLessonDTO(lesson))
		fmt.Println()
	}
}

func (h *Handler) createCourseLesson(c *Console) {
	var courseID domain.ID
	err := dto.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	var createLessonDTO dto.CreateLessonDTO
	err = dto.InputCreateLessonDTO(&createLessonDTO)
	if err != nil {
		ErrorResponse(err)
		return
	}

	var lesson domain.Lesson
	switch createLessonDTO.Type {
	case dto.LessonDTOTheory:
		if !createLessonDTO.Theory.Valid {
			ErrorResponse(BadRequestError)
			return
		}
		lesson, err = h.lessonService.CreateTheoryLesson(context.Background(),
			courseID, port.CreateTheoryParam{
				Title:  createLessonDTO.Title,
				Score:  int(createLessonDTO.Score.Int64),
				Theory: createLessonDTO.Theory.String,
			})
	case dto.LessonDTOVideo:
		if !createLessonDTO.VideoUrl.Valid {
			ErrorResponse(BadRequestError)
			return
		}
		lesson, err = h.lessonService.CreateVideoLesson(context.Background(),
			courseID, port.CreateVideoParam{
				Title:    createLessonDTO.Title,
				Score:    int(createLessonDTO.Score.Int64),
				VideoUrl: createLessonDTO.VideoUrl.String,
			})
	case dto.LessonDTOPractice:
		if createLessonDTO.Tests == nil {
			ErrorResponse(BadRequestError)
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
		lesson, err = h.lessonService.CreatePracticeLesson(context.Background(),
			courseID, port.CreatePracticeParam{
				Title: createLessonDTO.Title,
				Score: int(createLessonDTO.Score.Int64),
				Tests: tests,
			})
	default:
		ErrorResponse(BadRequestError)
		return
	}
	if err != nil {
		ErrorResponse(err)
		return
	}

	lessonDTO := dto.NewLessonDTO(lesson)
	dto.PrintLessonDTO(lessonDTO)
}

//	func (h *Handler) updateCourseLesson(c *Console) {
//		lessonID, err := getIdFromPath(context, "lesson_id")
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		var updateLessonDTO dto.UpdateLessonDTO
//		err = context.ShouldBindJSON(&updateLessonDTO)
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		lesson, err := h.lessonService.FindByID(context.Request.Context(), lessonID)
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		var updatedLesson domain.Lesson
//		switch lesson.Type {
//		case domain.TheoryLesson:
//			if !updateLessonDTO.Theory.Valid {
//				ErrorResponse(context, BadRequestError)
//				return
//			}
//			updatedLesson, err = h.lessonService.UpdateTheoryLesson(context.Request.Context(),
//				lessonID, port.UpdateTheoryParam{
//					Title:  updateLessonDTO.Title,
//					Score:  updateLessonDTO.Score,
//					Theory: updateLessonDTO.Theory,
//				})
//		case domain.VideoLesson:
//			if !updateLessonDTO.VideoUrl.Valid {
//				ErrorResponse(context, BadRequestError)
//				return
//			}
//			updatedLesson, err = h.lessonService.UpdateVideoLesson(context.Request.Context(),
//				lessonID, port.UpdateVideoParam{
//					Title:    updateLessonDTO.Title,
//					Score:    updateLessonDTO.Score,
//					VideoUrl: updateLessonDTO.VideoUrl,
//				})
//		case domain.PracticeLesson:
//			if updateLessonDTO.Tests == nil {
//				ErrorResponse(context, BadRequestError)
//				return
//			}
//
//			tests := make([]port.UpdateTestParam, len(updateLessonDTO.Tests))
//			for i, test := range updateLessonDTO.Tests {
//				tests[i] = port.UpdateTestParam{
//					Task:    test.Task,
//					Options: test.Options,
//					Answer:  test.Answer,
//					Level:   int(test.Level.Int64),
//					Score:   int(test.Score.Int64),
//				}
//			}
//			updatedLesson, err = h.lessonService.UpdatePracticeLesson(context.Request.Context(),
//				lessonID, port.UpdatePracticeParam{
//					Title: updateLessonDTO.Title,
//					Score: updateLessonDTO.Score,
//					Tests: tests,
//				})
//		default:
//			ErrorResponse(context, BadRequestError)
//			return
//		}
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		lessonDTO := dto.NewLessonDTO(updatedLesson)
//		CreatedResponse(context, lessonDTO)
//	}
//
//	func (h *Handler) deleteCourseLesson(c *Console) {
//		lessonID, err := getIdFromPath(context, "lesson_id")
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		err = h.lessonService.Delete(context.Request.Context(), lessonID)
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		SuccessResponse(context, "lesson successfully deleted")
//	}
//
//	func (h *Handler) addCourseReview(c *Console) {
//		userID, err := getIdFromRequestContext(context)
//		if err != nil {
//			ErrorResponse(context, UnauthorizedError)
//			return
//		}
//
//		courseID, err := getIdFromPath(context, "id")
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		var createReviewDTO dto.CreateReviewDTO
//		err = context.ShouldBindJSON(&createReviewDTO)
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		review, err := h.reviewService.CreateCourseReview(context, courseID, userID,
//			port.CreateReviewParam{Text: createReviewDTO.Text})
//
//		reviewDTO := dto.NewReviewDTO(review)
//		SuccessResponse(context, reviewDTO)
//	}
//
//	func (h *Handler) findCourseReviews(c *Console) {
//		courseID, err := getIdFromPath(context, "id")
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		reviews, err := h.reviewService.FindCourseReviews(context.Request.Context(), courseID)
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		reviewDTOs := make([]dto.ReviewDTO, len(reviews))
//		for i, review := range reviews {
//			reviewDTOs[i] = dto.NewReviewDTO(review)
//		}
//
//		SuccessResponse(context, reviewDTOs)
//	}
//
//	func (h *Handler) findLessonStat(c *Console) {
//		lessonID, err := getIdFromPath(context, "lesson_id")
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		userID, err := getIdFromRequestContext(context)
//		if err != nil {
//			ErrorResponse(context, UnauthorizedError)
//			return
//		}
//
//		stat, err := h.statService.FindLessonStat(context.Request.Context(), userID, lessonID)
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		statDTO := dto.NewLessonStatDTO(stat)
//		SuccessResponse(context, statDTO)
//	}
//
//	func (h *Handler) passCourseLesson(c *Console) {
//		lessonID, err := getIdFromPath(context, "lesson_id")
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		userID, err := getIdFromRequestContext(context)
//		if err != nil {
//			ErrorResponse(context, UnauthorizedError)
//			return
//		}
//
//		var passLessonDTO dto.PassLessonDTO
//		err = context.ShouldBindJSON(&passLessonDTO)
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		lesson, err := h.lessonService.FindByID(context, lessonID)
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		switch lesson.Type {
//		case domain.TheoryLesson:
//			fallthrough
//		case domain.VideoLesson:
//			err = h.statService.UpdateLessonStat(context.Request.Context(), userID,
//				lessonID, port.UpdateLessonStatParam{
//					Score:     null.IntFrom(int64(lesson.Score)),
//					TestStats: nil,
//				})
//		case domain.PracticeLesson:
//			if len(passLessonDTO.PassTests) == 0 {
//				ErrorResponse(context, BadRequestError)
//				return
//			}
//
//			testStats := make([]port.UpdateTestStatParam, len(passLessonDTO.PassTests))
//			for i, passTest := range passLessonDTO.PassTests {
//				for j, test := range lesson.Tests {
//					if passTest.TestID == test.ID.String() {
//						var newScore int
//						if passTest.Answer == test.Answer {
//							newScore = lesson.Tests[j].Score
//						}
//
//						testStats[i] = port.UpdateTestStatParam{
//							TestID: test.ID,
//							Score:  newScore,
//						}
//					}
//				}
//			}
//
//			err = h.statService.UpdateLessonStat(context.Request.Context(), userID,
//				lessonID, port.UpdateLessonStatParam{
//					Score:     null.IntFrom(int64(lesson.Score)),
//					TestStats: testStats,
//				})
//		}
//
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		SuccessResponse(context, "successfully passed lesson")
//	}
//
//	func (h *Handler) verifyCourseWriteAccess(context *gin.Context) {
//		courseID, err := getIdFromPath(context, "id")
//		if err != nil {
//			ErrorResponse(context, err)
//			return
//		}
//
//		if !h.checkCurrentUserIsCourseTeacher(context, courseID) {
//			ErrorResponse(context, ForbiddenError)
//			return
//		}
//	}

func (h *Handler) verifyCourseReadAccess(c *Console, courseID domain.ID) bool {
	return h.checkCurrentUserIsCourseOwner(c, courseID) ||
		h.checkCurrentUserIsCourseTeacher(c, courseID) ||
		h.checkCurrentUserIsCourseStudent(c, courseID)
}

func (h *Handler) checkCurrentUserIsCourseStudent(c *Console, courseID domain.ID) bool {
	err := h.verifyAuth(c)
	if err != nil {
		return false
	}
	userID := *c.UserID

	isStudent, err := h.courseService.IsCourseStudent(context.Background(), userID, courseID)
	if err != nil {
		return false
	}

	return isStudent
}

func (h *Handler) checkCurrentUserIsCourseTeacher(c *Console, courseID domain.ID) bool {
	err := h.verifyAuth(c)
	if err != nil {
		return false
	}
	userID := *c.UserID

	isTeacher, err := h.courseService.IsCourseTeacher(context.Background(), userID, courseID)
	if err != nil {
		ErrorResponse(err)
		return false
	}

	return isTeacher
}

func (h *Handler) checkCurrentUserIsCourseOwner(c *Console, courseID domain.ID) bool {
	err := h.verifyAuth(c)
	if err != nil {
		return false
	}

	course, err := h.courseService.FindByID(context.Background(), courseID)
	if err != nil {
		ErrorResponse(err)
		return false
	}

	return h.checkCurrentUserIsSchoolOwner(c, course.SchoolID)
}
