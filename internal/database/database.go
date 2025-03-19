package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
	"github.com/yummynaan/kit-enrollment-helper/internal/domain/model"
)

const (
	tableCourses = "courses"
)

var (
	host     string
	port     string
	username string
	password string
	dbname   string
)

type Database struct {
	sess *dbr.Session
}

func NewDatabase() (*Database, error) {
	if host = os.Getenv("DB_HOST"); host == "" {
		host = "localhost"
	}
	if port = os.Getenv("DB_PORT"); port == "" {
		port = "3308"
	}
	if username = os.Getenv("DB_USER"); username == "" {
		username = "root"
	}
	if password = os.Getenv("DB_PASSWORD"); password == "" {
		password = "root"
	}
	if dbname = os.Getenv("DB_NAME"); dbname == "" {
		dbname = "kit_enrollment_helper"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbname)
	conn, err := dbr.Open("mysql", dsn, nil)
	if err != nil {
		log.Println("failed at dbr.Open()")
		return nil, err
	}

	sess := conn.NewSession(nil)

	return &Database{sess: sess}, nil
}

func (db *Database) CreateUser(user model.User) error {
	return nil
}

func (db *Database) GetUserByEmail(email string) (model.User, error) {
	return model.User{}, nil
}

func (db *Database) BulkInsertCourses(courses model.Courses) (int64, error) {
	tx, err := db.sess.Begin()
	if err != nil {
		return 0, err
	}

	cols := []string{"timetable_id", "title", "class", "type", "credits", "instructors", "year", "semester", "day"}
	stmt := tx.InsertInto(tableCourses).Columns(cols...)
	for _, v := range courses {
		c := course{
			TimetableID: v.TimetableID,
			Class:       v.Class,
			Type:        v.Type,
			Credits:     v.Credits,
			Instructors: v.Instructors,
			Title:       v.Title,
			Year:        v.Year,
			Semester:    v.Semester,
			Day:         v.Day,
		}
		if v.TimetableID == nil {
			c.TimetableID = nil
		}
		stmt = stmt.Record(&c)
	}
	result, err := stmt.Exec()
	if err != nil {
		return 0, err
	}
	n, _ := result.RowsAffected()

	return n, tx.Commit()
}
