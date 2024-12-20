package console

import (
	"bufio"
	"context"
	"fmt"
	"github.com/guregu/null"
	dto2 "github.com/paw1a/eschool/internal/adapter/delivery/console/dto"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"io"
	"net/http"
	"os"
	"strings"
)

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
		dto2.PrintCourseDTO(dto2.NewCourseDTO(course))
		fmt.Println()
	}
}

func (h *Handler) FindCourseByID(c *Console) {
	var courseID domain.ID
	err := dto2.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	course, err := h.courseService.FindByID(context.Background(), courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	courseDTO := dto2.NewCourseDTO(course)
	dto2.PrintCourseDTO(courseDTO)
}

func (h *Handler) FindLessonByID(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}

	var lessonID domain.ID
	err = dto2.InputID(&lessonID, "lesson")
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

	lessonDTO := dto2.NewLessonDTO(lesson)
	dto2.PrintLessonDTO(lessonDTO)
}

func (h *Handler) FindCourseTeachers(c *Console) {
	var courseID domain.ID
	err := dto2.InputID(&courseID, "course")
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
		dto2.PrintUserDTO(dto2.NewUserDTO(teacher))
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
	err = dto2.InputID(&courseID, "course")
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
	err = dto2.InputID(&teacherID, "teacher")
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

func (h *Handler) PublishCourse(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}

	var courseID domain.ID
	err = dto2.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	if !h.verifyCourseWriteAccess(c, courseID) {
		fmt.Println("you are not a course teacher")
		return
	}

	errList := h.courseService.ConfirmDraftCourse(context.Background(), courseID)
	if errList != nil {
		ErrorResponse(errList[0])
		return
	}

	err = h.courseService.PublishReadyCourse(context.Background(), courseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("course successfully published")
}

func (h *Handler) FindCourseLessons(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}

	var courseID domain.ID
	err = dto2.InputID(&courseID, "course")
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
		dto2.PrintLessonDTO(dto2.NewLessonDTO(lesson))
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
	err = dto2.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	if !h.verifyCourseWriteAccess(c, courseID) {
		ErrorResponse(ForbiddenError)
		return
	}

	var createLessonDTO dto2.CreateLessonDTO
	err = dto2.InputCreateLessonDTO(&createLessonDTO)
	if err != nil {
		ErrorResponse(err)
		return
	}

	var lesson domain.Lesson
	switch createLessonDTO.Type {
	case dto2.LessonDTOTheory:
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
	case dto2.LessonDTOVideo:
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
	case dto2.LessonDTOPractice:
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

	lessonDTO := dto2.NewLessonDTO(lesson)
	dto2.PrintLessonDTO(lessonDTO)
}

func (h *Handler) AddCourseReview(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}
	userID := *c.UserID

	var courseID domain.ID
	err = dto2.InputID(&courseID, "course")
	if err != nil {
		ErrorResponse(err)
		return
	}

	var createReviewDTO dto2.CreateReviewDTO
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Review text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	createReviewDTO.Text = text

	review, err := h.reviewService.CreateCourseReview(context.Background(), courseID, userID,
		port.CreateReviewParam{Text: createReviewDTO.Text})

	reviewDTO := dto2.NewReviewDTO(review)
	dto2.PrintReviewDTO(reviewDTO)
}

func (h *Handler) FindCourseReviews(c *Console) {
	var courseID domain.ID
	err := dto2.InputID(&courseID, "course")
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
		dto2.PrintReviewDTO(dto2.NewReviewDTO(review))
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
	err = dto2.InputID(&lessonID, "lesson")
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
		fmt.Println("you are not a student of this course")
		return
	}

	stat, err := h.statService.FindLessonStat(context.Background(), userID, lessonID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	statDTO := dto2.NewLessonStatDTO(stat)
	dto2.PrintLessonStatDTO(statDTO)
}

func (h *Handler) PassCourseLesson(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		ErrorResponse(UnauthorizedError)
		return
	}
	userID := *c.UserID

	var lessonID domain.ID
	err = dto2.InputID(&lessonID, "lesson")
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
		fmt.Println("you are not a student of this course")
		return
	}

	course, err := h.courseService.FindByID(context.Background(), lesson.CourseID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if course.Status != domain.CoursePublished {
		fmt.Println("course is not published yet")
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
