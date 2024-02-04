package main

import (
	"fmt"
	"net/http"

	"cdn-admin/internal/handlers"
	"cdn-admin/pkg/logger"

	"github.com/gorilla/mux"
)

func main() {
	logger.InitLogger()

	r := mux.NewRouter()
	r.HandleFunc("/upload", handlers.UploadFileHandler).Methods("POST")
	r.HandleFunc("/download/{filename}", handlers.DownloadFileHandler).Methods("GET")
	r.HandleFunc("/delete/{filename}", handlers.DeleteFileHandler).Methods("DELETE")

	http.Handle("/", r)

	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		logger.ErrorLogger.Println("Error starting server:", err)
	}
}
