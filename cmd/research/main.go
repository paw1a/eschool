package main

import (
	"context"
	"database/sql"
	"fmt"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	"github.com/paw1a/eschool/internal/app/config"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/pkg/database/postgres"
	"github.com/paw1a/eschool/pkg/logging"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "paw1a"
	password = "0023"
	dbname   = "eschool"
)

func executeSQLFile(db *sql.DB, filepath string) error {
	sqlBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	sqlStr := string(sqlBytes)
	_, err = db.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("failed to execute SQL file: %w", err)
	}

	return nil
}

func clearTables(db *sql.DB) error {
	tables := []string{"public.test", "public.lesson", "public.course", "public.school", "public.user"}

	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", table)
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to truncate table %s: %w", table, err)
		}
	}

	return nil
}

func runPythonScript(users, schools, courses, lessons, tests int, outputFile string) error {
	cmd := exec.Command("python3", "cmd/research/main.py",
		"--users", strconv.Itoa(users),
		"--schools", strconv.Itoa(schools),
		"--courses", strconv.Itoa(courses),
		"--lessons", strconv.Itoa(lessons),
		"--tests", strconv.Itoa(tests),
		"--output", outputFile)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run Python script: %w", err)
	}

	return nil
}

func measureTriggerExecutionTime(db *sql.DB) (error, time.Duration) {
	start := time.Now()
	query := "UPDATE public.course SET status = 'draft';"
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to update course status: %w", err), 0
	}
	elapsed := time.Since(start)
	return nil, elapsed
}

func measureAppExecutionTime(courseService *service.CourseService) (error, time.Duration) {
	courses, err := courseService.FindAll(context.Background())
	if err != nil {
		log.Fatalf("failed to find courses: %v", err)
		return err, 0
	}

	start := time.Now()
	for _, course := range courses {
		errors := courseService.ConfirmDraftCourse(context.Background(), course.ID)
		if errors != nil {
			log.Fatal(fmt.Sprintf("failed to confirm draft course: %v", errors))
			return err, 0
		}
	}
	elapsed := time.Since(start)

	return nil, elapsed
}

func generateData(db *sql.DB, users, schools, courses, lessons, tests int) error {
	outputFile := "output.sql"
	err := runPythonScript(users, schools, courses, lessons, tests, outputFile)
	if err != nil {
		log.Fatalf("failed to run Python script: %v", err)
		return err
	}

	err = clearTables(db)
	if err != nil {
		log.Fatalf("failed to clear tables: %v", err)
		return err
	}

	err = executeSQLFile(db, outputFile)
	if err != nil {
		log.Fatalf("failed to execute SQL file: %v", err)
		return err
	}

	return os.RemoveAll(outputFile)
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", host, port, user, password, dbname)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer conn.Close()

	cfg := config.GetConfig()
	logger, err := logging.NewLogger(&cfg.Logging)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	root, err := postgres.NewPostgresDB(&cfg.PostgresRoot, logger)
	if err != nil {
		log.Fatal("failed to create postgres root connection")
	}

	guest, err := postgres.NewPostgresDB(&cfg.PostgresGuest, logger)
	if err != nil {
		log.Fatal("failed to create postgres guest connection")
	}

	authenticated, err := postgres.NewPostgresDB(&cfg.PostgresAuthenticated, logger)
	if err != nil {
		log.Fatal("failed to create postgres authenticated connection")
	}

	db := postgres.DB{
		Root:          root,
		Guest:         guest,
		Authenticated: authenticated,
	}

	courseRepo := repository.NewCourseRepo(&db)
	lessonRepo := repository.NewLessonRepo(&db)
	schoolRepo := repository.NewSchoolRepo(&db)
	statRepo := repository.NewStatRepo(&db)
	courseService := service.NewCourseService(courseRepo, lessonRepo,
		schoolRepo, statRepo, logger)

	file, err := os.Create("cmd/research/data.txt")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	for i := 20; i <= 200; i += 10 {
		err = generateData(conn, 10, 10, 10, i, 20)
		if err != nil {
			log.Fatalf("failed to generate data: %v", err)
		}

		err, delta1 := measureAppExecutionTime(courseService)
		if err != nil {
			log.Fatalf("failed to measure trigger execution time: %v", err)
		}

		err, delta2 := measureTriggerExecutionTime(conn)
		if err != nil {
			log.Fatalf("failed to measure app execution time: %v", err)
		}

		_, err = file.WriteString(fmt.Sprintf("%d %d %d\n",
			i, delta1.Milliseconds(), delta2.Milliseconds()))
		if err != nil {
			log.Fatalf("failed to write to file: %v", err)
		}
	}
}
