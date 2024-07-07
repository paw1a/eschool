package console

import (
	"context"
	"fmt"
	"github.com/paw1a/eschool/internal/core/domain"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

type Console struct {
	Handler *Handler
	Routes  map[Option]func(*Console)
	UserID  *domain.ID
	Logger  *zap.Logger
}

type Option int

const (
	exit Option = iota
	menu

	signIn
	signUp
	logout
	getAllUsers
	getUserAccount
	updateUserAccount
	findUserCourses
	addUserFreeCourse

	findAllCourses
	createSchoolCourse
	updateSchoolCourse
	deleteSchoolCourse
	findCourseLessons
	createCourseLesson
	findCourseTeachers
	addCourseTeacher
	publishCourse

	findAllSchools
	createSchool
	updateSchool
	findSchoolCourses
	findSchoolTeachers
	addSchoolTeacher

	passCourseLesson
	findLessonStat

	findCourseReviews
	addCourseReview

	getCourseCertificate
	createCourseCertificate
)

func NewConsole(lc fx.Lifecycle, handler *Handler, logger *zap.Logger) *Console {
	c := &Console{
		Handler: handler,
		Logger:  logger,
	}
	c.InitRoutes()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("console interface started")
				log.Fatal(c.Start())
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
	return c
}

func (c *Console) InitRoutes() {
	c.Routes = map[Option]func(*Console){
		exit: func(c *Console) { os.Exit(0) },
		menu: func(c *Console) { c.PrintMenu() },

		signIn: c.Handler.UserSignIn,
		signUp: c.Handler.UserSignUp,
		logout: c.Handler.UserLogout,

		getAllUsers:       c.Handler.GetAllUsers,
		getUserAccount:    c.Handler.GetUserAccount,
		updateUserAccount: c.Handler.UpdateUser,
		findUserCourses:   c.Handler.FindUserCourses,
		addUserFreeCourse: c.Handler.AddUserFreeCourse,

		findAllCourses:     c.Handler.FindAllCourses,
		createSchoolCourse: c.Handler.CreateSchoolCourse,
		updateSchoolCourse: c.Handler.UpdateSchoolCourse,
		deleteSchoolCourse: c.Handler.DeleteSchoolCourse,
		findCourseLessons:  c.Handler.FindCourseLessons,
		createCourseLesson: c.Handler.CreateCourseLesson,
		findCourseTeachers: c.Handler.FindCourseTeachers,
		addCourseTeacher:   c.Handler.AddCourseTeacher,
		publishCourse:      c.Handler.PublishCourse,

		findAllSchools:     c.Handler.FindAllSchools,
		createSchool:       c.Handler.CreateSchool,
		updateSchool:       c.Handler.UpdateSchool,
		findSchoolCourses:  c.Handler.FindSchoolCourses,
		findSchoolTeachers: c.Handler.FindSchoolTeachers,
		addSchoolTeacher:   c.Handler.AddSchoolTeacher,

		passCourseLesson: c.Handler.PassCourseLesson,
		findLessonStat:   c.Handler.FindLessonStat,

		findCourseReviews: c.Handler.FindCourseReviews,
		addCourseReview:   c.Handler.AddCourseReview,

		getCourseCertificate:    c.Handler.GetCourseCertificate,
		createCourseCertificate: c.Handler.CreateCourseCertificate,
	}
}

func (c *Console) Start() error {
	time.Sleep(1 * time.Second)
	for {
		c.PrintMenu()

		var option Option
		fmt.Print("Choose menu option: ")
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			fmt.Println("Invalid menu option")
			continue
		}
		fmt.Println()

		handleFunc, ok := c.Routes[option]
		if !ok {
			fmt.Println("Invalid menu option")
			continue
		}
		handleFunc(c)
	}
}

func (c *Console) PrintMenu() {
	fmt.Println()
	fmt.Println("--------------------------------")
	fmt.Println("Menu")
	fmt.Println("0  Exit")
	fmt.Println("1  Print menu")
	fmt.Println("2  Sign In")
	fmt.Println("3  Sign Up")
	fmt.Println("4  Logout")
	fmt.Println("5  Get all users")
	fmt.Println("6  Get user account")
	fmt.Println("7  Update user account")
	fmt.Println("8  Get user purchased courses")
	fmt.Println("9  Add free course to user library")

	fmt.Println("10 Get all courses")
	fmt.Println("11 Create course")
	fmt.Println("12 Update course")
	fmt.Println("13 Delete course")
	fmt.Println("14 Get course lessons")
	fmt.Println("15 Create course lesson")
	fmt.Println("16 Get course teachers")
	fmt.Println("17 Add course teacher")
	fmt.Println("18 Publish course")

	fmt.Println("19 Get all schools")
	fmt.Println("20 Create school")
	fmt.Println("21 Update school")
	fmt.Println("22 Get school courses")
	fmt.Println("23 Get school teachers")
	fmt.Println("24 Add school teacher")

	fmt.Println("25 Pass course lesson")
	fmt.Println("26 Get lesson progress")

	fmt.Println("27 Find course reviews")
	fmt.Println("28 Add course review")

	fmt.Println("29 Get course certificate")
	fmt.Println("30 Generate course certificate")
	fmt.Println("--------------------------------")
}
