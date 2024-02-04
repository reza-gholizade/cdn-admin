package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"cdn-admin/pkg/logger"
	"github.com/gorilla/mux"
)

const uploadDirectory = "./uploads/"

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		logger.ErrorLogger.Println("Error retrieving the file:", err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := handler.Filename
	dstPath := filepath.Join(uploadDirectory, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		logger.ErrorLogger.Println("Error creating the file:", err)
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	io.Copy(dst, file)

	absolutePath, err := filepath.Abs(dstPath)
	if err != nil {
		logger.ErrorLogger.Println("Error getting absolute path:", err)
		http.Error(w, "Error getting absolute path", http.StatusInternalServerError)
		return
	}

	logger.InfoLogger.Println(fmt.Sprintf("File %s uploaded successfully to %s", filename, absolutePath))

	response := fmt.Sprintf("File %s uploaded successfully. Absolute path: %s", filename, absolutePath)
	fmt.Fprintf(w, response)
}


func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	filePath := filepath.Join(uploadDirectory, filename)

	http.ServeFile(w, r, filePath)
	logger.InfoLogger.Println(fmt.Sprintf("File %s downloaded successfully!", filename))
}

func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	filePath := filepath.Join(uploadDirectory, filename)

	err := os.Remove(filePath)
	if err != nil {
		logger.ErrorLogger.Println("Error deleting the file:", err)
		http.Error(w, "Error deleting the file", http.StatusInternalServerError)
		return
	}

	logger.InfoLogger.Println(fmt.Sprintf("File %s deleted successfully!", filename))
	fmt.Fprintf(w, "File %s deleted successfully!", filename)
}
