package model

type Course struct {
	ID          int64
	TimetableID *int64
	Class       *string
	Type        string
	Credits     int
	Instructors string
	Title       string
	Year        int
	Semester    string
	Day         *string
}

type Courses []Course
