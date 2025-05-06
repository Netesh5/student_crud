package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/netesh5/student_crud/internal/config"
	student "github.com/netesh5/student_crud/internal/http/handlers"
	"github.com/netesh5/student_crud/internal/storage/sqlite"
)

func main() {
	config := config.MustLoad()
	println(config.Env)

	storage, err := sqlite.New(config)
	if err != nil {
		println("Error in creating storage", err.Error())
		return
	}
	slog.Info("Storage initialized successfully", slog.String("storage_path", config.StoragePath), slog.String("env", config.Env))

	router := mux.NewRouter()
	router.HandleFunc("/", initialPage).Methods("GET")
	router.HandleFunc("/students", student.StudentHandler(storage)).Methods("GET")
	router.HandleFunc("/students", student.StudentHandler(storage)).Methods("POST")

	server := http.Server{
		Addr:    config.Server.Address,
		Handler: router,
	}
	println("Server is running on port", config.Server.Address)

	done := make(chan os.Signal, 1)

	// Listen for interrupt/termination signals
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed", "error", err.Error())
		}
	}()

	<-done // Block until a signal is received
	slog.Info("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err.Error())
	} else {
		slog.Info("Server exited gracefully")
	}

}

func initialPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Student CRUD API"))
}
