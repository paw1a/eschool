package console

import (
	"bufio"
	"context"
	"fmt"
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/adapter/delivery/console/dto"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"io"
	"net/http"
	"os"
	"strings"
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
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}

	var lessonID domain.ID
	err = dto.InputID(&lessonID, "lesson")
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

func (h *Handler) CreateCourseLesson(c *Console) {
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

	if !h.verifyCourseWriteAccess(c, courseID) {
		ErrorResponse(ForbiddenError)
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
func (h *Handler) AddCourseReview(c *Console) {
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

	var createReviewDTO dto.CreateReviewDTO
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Review text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	createReviewDTO.Text = text

	review, err := h.reviewService.CreateCourseReview(context.Background(), courseID, userID,
		port.CreateReviewParam{Text: createReviewDTO.Text})

	reviewDTO := dto.NewReviewDTO(review)
	dto.PrintReviewDTO(reviewDTO)
}

func (h *Handler) FindCourseReviews(c *Console) {
	var courseID domain.ID
	err := dto.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	reviews, err := h.reviewService.FindCourseReviews(context.Background(), courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(reviews) == 0 {
		fmt.Println("no reviews")
		return
	}

	for _, review := range reviews {
		dto.PrintReviewDTO(dto.NewReviewDTO(review))
		fmt.Println()
	}
}

func (h *Handler) FindLessonStat(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}
	userID := *c.UserID

	var lessonID domain.ID
	err = dto.InputID(&lessonID, "lesson")
	if err != nil {
		ErrorResponse(err)
		return
	}

	stat, err := h.statService.FindLessonStat(context.Background(), userID, lessonID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	statDTO := dto.NewLessonStatDTO(stat)
	dto.PrintLessonStatDTO(statDTO)
}

func (h *Handler) PassCourseLesson(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}
	userID := *c.UserID

	var lessonID domain.ID
	err = dto.InputID(&lessonID, "lesson")
	if err != nil {
		ErrorResponse(err)
		return
	}

	lesson, err := h.lessonService.FindByID(context.Background(), lessonID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	switch lesson.Type {
	case domain.TheoryLesson:
		fallthrough
	case domain.VideoLesson:
		err = h.statService.UpdateLessonStat(context.Background(), userID,
			lessonID, port.UpdateLessonStatParam{
				Score:     null.IntFrom(int64(lesson.Score)),
				TestStats: nil,
			})
	case domain.PracticeLesson:
		testStats := make([]port.UpdateTestStatParam, len(lesson.Tests))
		for i, test := range lesson.Tests {
			fmt.Printf("Test #%d\n", i+1)
			response, err := http.Get(test.TaskUrl)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer response.Body.Close()
			_, err = io.Copy(os.Stdout, response.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println()
			fmt.Println("Available answers:")
			for _, option := range test.Options {
				fmt.Println(option)
			}

			fmt.Print("Choose correct answer: ")
			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)
			fmt.Println()

			var newScore int
			if answer == test.Answer {
				newScore = lesson.Tests[i].Score
			}

			testStats[i] = port.UpdateTestStatParam{
				TestID: test.ID,
				Score:  newScore,
			}
		}

		err = h.statService.UpdateLessonStat(context.Background(), userID,
			lessonID, port.UpdateLessonStatParam{
				Score:     null.IntFrom(int64(lesson.Score)),
				TestStats: testStats,
			})
	}

	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("successfully passed lesson")
}

func (h *Handler) GetCourseCertificate(c *Console) {
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

	certificate, err := h.certificateService.FindCourseCertificate(context.Background(),
		courseID, userID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	certificateDTO := dto.NewCertificateDTO(certificate)
	dto.PrintCertificateDTO(certificateDTO)
}

func (h *Handler) CreateCourseCertificate(c *Console) {
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

	certificate, err := h.certificateService.FindCourseCertificate(context.Background(),
		courseID, userID)
	if err == nil {
		fmt.Println("Certificate for this course is already exists")
		return
	}

	certificate, err = h.certificateService.CreateCourseCertificate(context.Background(),
		userID, courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	certificateDTO := dto.NewCertificateDTO(certificate)
	dto.PrintCertificateDTO(certificateDTO)
}

func (h *Handler) verifyCourseWriteAccess(c *Console, courseID domain.ID) bool {
	return h.checkCurrentUserIsCourseTeacher(c, courseID)
}

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
