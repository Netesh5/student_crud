package student

import (
	"encoding/json"
	"net/http"

	types "github.com/netesh5/student_crud/internal/type"
	"github.com/netesh5/student_crud/internal/utils/response"
)

func StudentHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			response.ToJson(w, http.StatusBadRequest, err.Error())
			return
		}
		w.Write([]byte("Welcome to the Student CRUD API"))
	}

}
