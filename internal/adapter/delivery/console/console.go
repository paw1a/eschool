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

	getUserAccount
	updateUserAccount
	findUserCourses
	addUserFreeCourse

	findAllCourses
	findCourseLessons
	findCourseTeachers
	addCourseTeacher
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

		getUserAccount:    c.Handler.GetUserAccount,
		updateUserAccount: c.Handler.UpdateUser,
		findUserCourses:   c.Handler.FindUserCourses,
		addUserFreeCourse: c.Handler.AddUserFreeCourse,

		findAllCourses:     c.Handler.FindAllCourses,
		findCourseLessons:  c.Handler.FindCourseLessons,
		findCourseTeachers: c.Handler.FindCourseTeachers,
		addCourseTeacher:   c.Handler.AddCourseTeacher,
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
	fmt.Println("5  Get user account")
	fmt.Println("6  Update user account")
	fmt.Println("7  Get user purchased courses")
	fmt.Println("8  Add free course to user library")
	fmt.Println("9  Get all courses")
	fmt.Println("10 Get course lessons")
	fmt.Println("11 Get course teachers")
	fmt.Println("12 Add course teacher")
	fmt.Println("--------------------------------")
}
