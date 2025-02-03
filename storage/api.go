package storage

import (
	"fmt"
	"io"
	"net/http"
)

var storage Storage = Storage{path: "/Users/sergei/code/sergio-fry/file-storage/data"}

func getFile(w http.ResponseWriter, r *http.Request) {
	var result = storage.read(r.URL.Query().Get("name"))
	defer result.Close()
	io.Copy(w, result)
}

func main() {
	http.HandleFunc("/files", getFile)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
