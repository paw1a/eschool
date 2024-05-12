package console

import (
	"context"
	"fmt"
	"github.com/paw1a/eschool/internal/app/config"
	"github.com/paw1a/eschool/internal/core/domain"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"os"
	"time"
)

type Console struct {
	Cfg     *config.Config
	Handler *Handler
	Routes  map[Option]func(*Console)
	UserID  *domain.ID
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

func NewConsole(lc fx.Lifecycle, cfg *config.Config, handler *Handler) *Console {
	c := &Console{
		Cfg:     cfg,
		Handler: handler,
	}
	c.InitRoutes()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Infof("Console interface started")
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

	fmt.Println("10  Get all courses")
	fmt.Println("11 Create course")
	fmt.Println("12 Update course")
	fmt.Println("13 Delete course")
	fmt.Println("14 Get course lessons")
	fmt.Println("15 Create course lesson")
	fmt.Println("16 Get course teachers")
	fmt.Println("17 Add course teacher")

	fmt.Println("18 Get all schools")
	fmt.Println("19 Create school")
	fmt.Println("20 Update school")
	fmt.Println("21 Get school courses")
	fmt.Println("22 Get school teachers")
	fmt.Println("23 Add school teacher")

	fmt.Println("24 Pass course lesson")
	fmt.Println("25 Get lesson progress")

	fmt.Println("26 Find course reviews")
	fmt.Println("27 Add course review")

	fmt.Println("28 Get course certificate")
	fmt.Println("29 Generate course certificate")
	fmt.Println("--------------------------------")
}
