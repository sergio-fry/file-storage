package storage

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
)

var globalStorage = Storage{path: "/Users/sergei/code/sergio-fry/file-storage/data"}

func getFile(w http.ResponseWriter, r *http.Request) {
	var result = globalStorage.read(r.URL.Query().Get("name"))
	defer result.Close()
	io.Copy(w, result)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving the file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := filepath.Base(header.Filename)

	globalStorage.write(filename, file)

}

func main() {
	http.HandleFunc("/files", getFile)
	http.HandleFunc("/upload", uploadFile)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
