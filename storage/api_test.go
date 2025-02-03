package storage

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetFiles(t *testing.T) {
	globalStorage.write("foo.txt", strings.NewReader("bar"))

	req, err := http.NewRequest("GET", "/files?name=foo.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	getFile(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "bar"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUploadFile(t *testing.T) {
	var file = strings.NewReader("bar")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "foo.txt")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing writer:", err)
		return
	}

	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	rr := httptest.NewRecorder()
	uploadFile(rr, req)

	var result []byte
	result, err = io.ReadAll(globalStorage.read("foo.txt"))

	if string(result) != "bar" {
		t.Errorf("handler returned unexpected body: got %v want %v", string(result), "bar")
	}
}
