package model

type User struct {
	ID             int64
	Name           string
	Email          string
	Password       string
	EnrollmentYear int
	Department     Department
}
