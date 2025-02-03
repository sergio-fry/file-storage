package storage

import (
	"io"
	"strings"
	"testing"
)

func TestWrite(t *testing.T) {
	var storage = Storage{path: "/Users/sergei/code/sergio-fry/file-storage/data"}

	storage.write("foo", strings.NewReader("bar"))

	var result, err = io.ReadAll(storage.read("foo"))

	if err != nil {
		t.Errorf("Ошибка чтения: %v", err)
	}

	if string(result) != "bar" {
		t.Errorf("Ожидаемый текст: %s, но получен: %s", "bar", string(result))
	}
}
