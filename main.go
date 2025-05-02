package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/netesh5/student_crud/internal/config"
)

func main() {
	config := config.MustLoad()
	println(config.Env)

	router := mux.NewRouter()

	router.HandleFunc("/students", getStudents).Methods("GET")

	server := http.Server{
		Addr:    config.Server.Address,
		Handler: router,
	}
	println("Server is running on port", config.Server.Address)
	if err := server.ListenAndServe(); err != nil {
		panic(err)

	}

}

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
