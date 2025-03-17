package database

import (
	"github.com/yummynaan/kit-enrollment-helper/internal/domain/model"
)

type Repository interface {
	CreateUser(user model.User) error
	GetUserByEmail(email string) (model.User, error)
	BulkInsertCourses(courses model.Courses) (int64, error)
}
