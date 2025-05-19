package sqlite

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/netesh5/student_crud/internal/config"

	types "github.com/netesh5/student_crud/internal/type"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		age INTEGER NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(student types.Student) (int64, error) {
	// Check if email already exists
	var existingId int64
	err := s.Db.QueryRow(`SELECT id FROM students WHERE email = ?`, student.Email).Scan(&existingId)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if existingId != 0 {
		return 0, fmt.Errorf("email already exists")
	}

	stmt, err := s.Db.Prepare(`INSERT INTO students (name, email, age) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(student.Name, student.Email, student.Age)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare(`SELECT * FROM students WHERE id=? LIMIT 1`)
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age, &student.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with %s", strconv.Itoa(int(id)))
		}
		return types.Student{}, err
	}
	return student, nil
}

func (s *Sqlite) GetAllStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare(`SELECT * FROM students`)
	if err != nil {
		return []types.Student{}, err
	}

	var students []types.Student

	rows, err := stmt.Query()
	if err != nil {
		return []types.Student{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var student types.Student
		err = rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age, &student.CreatedAt)
		if err != nil {
			return []types.Student{}, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return []types.Student{}, err
	}
	return students, nil
}
