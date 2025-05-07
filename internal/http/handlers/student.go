package student

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/netesh5/student_crud/internal/storage"
	types "github.com/netesh5/student_crud/internal/type"
	"github.com/netesh5/student_crud/internal/utils/response"
)

func CreateStudent(s storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			response.ToJson(w, http.StatusBadRequest, response.ErrorResponse(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			response.ToJson(w, http.StatusBadRequest, response.ValidationErrorResponse(err.(validator.ValidationErrors)))
			return
		}

		id, err := s.CreateStudent(student)

		if err != nil {
			response.ToJson(w, http.StatusInternalServerError, response.ErrorResponse(err))
			return
		}

		response.ToJson(w, http.StatusOK, map[string]interface{}{
			"id":    strconv.Itoa(int(id)),
			"name":  student.Name,
			"email": student.Email,
			"age":   strconv.Itoa(student.Age),
			"meta": map[string]string{
				"status":  response.StatusOk,
				"message": "Student created successfully",
			},
		})

	}

}

func GetStudentById(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// id:=r.URL.Query().Get("id")
		id := r.URL.Path[len("/student/"):]

		if id == "" {
			response.ToJson(w, http.StatusBadRequest, response.ErrorResponse(errors.New("id is required")))
			return
		}

		parseId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.ToJson(w, http.StatusBadRequest, response.ErrorResponse(err))
			return
		}
		student, err := s.GetStudentById(parseId)
		if err != nil {
			response.ToJson(w, http.StatusInternalServerError, response.ErrorResponse(err))
			return
		}
		response.ToJson(w, http.StatusOK, student)
	}
}

func GetAllSutudents(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, err := s.GetAllStudents()
		if err != nil {
			response.ToJson(w, http.StatusInternalServerError, response.ErrorResponse(err))
			return
		}
		response.ToJson(w, http.StatusOK, students)

	}

}
