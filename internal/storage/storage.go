package storage

import (
	types "github.com/netesh5/student_crud/internal/type"
)

type Storage interface {
	CreateStudent(student types.Student) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
}
