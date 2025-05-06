package student

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/netesh5/student_crud/internal/storage"
	types "github.com/netesh5/student_crud/internal/type"
	"github.com/netesh5/student_crud/internal/utils/response"
)

func StudentHandler(storage *storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			response.ToJson(w, http.StatusBadRequest, response.ErrorResponse(err))
			return
		}
		// w.Write([]byte("Welcome to the Student CRUD API"))

		if err := validator.New().Struct(student); err != nil {
			response.ToJson(w, http.StatusBadRequest, response.ValidationErrorResponse(err.(validator.ValidationErrors)))
			return
		}

		response.ToJson(w, http.StatusOK, map[string]string{
			"message": "Student created successfully",
			"student": student.Name,
		})

	}

}
