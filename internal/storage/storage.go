package storage

import "github/sunil/prod-go-server/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64,error)
	GetStudentById(id int64) (types.Student, error) 
}