package model

type Course struct {
	ID          string
	TimetableID *string
	Class       *string
	Type        string
	Credits     int
	Title       string
	Semester    string
	Day         string
}
