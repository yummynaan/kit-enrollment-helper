package database

import "time"

type department struct {
	ID      int64  `db:"id"`
	Faculty string `db:"faculty"`
	Field   string `db:"field"`
	Program string `db:"program"`
}

type user struct {
	ID             int64     `db:"id"`
	Name           string    `db:"name"`
	Email          string    `db:"email"`
	PasswordHash   string    `db:"password_hash"`
	EnrollmentYear int       `db:"enrollment_year"`
	DepartmentID   int64     `db:"department_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type course struct {
	ID          int64   `db:"id"`
	TimetableID *string `db:"timetable_id"`
	Class       *string `db:"class"`
	Type        string  `db:"type"`
	Credits     int     `db:"credits"`
	Instructors string  `db:"instructors"`
	Title       string  `db:"title"`
	Year        string  `db:"year"`
	Semester    string  `db:"semester"`
	Day         *string `db:"day"`
}

func (c course) columns() []string {
	return []string{"id", "timetable_id", "class", "type", "credits", "instructors", "title", "year", "semester", "day"}
}
