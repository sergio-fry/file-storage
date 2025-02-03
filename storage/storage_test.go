package storage

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	filename := "test.txt"
	text := "Привет, мир!"

	writeToFile(filename, text)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Ошибка чтения файла: %v", err)
	}

	if string(data) != text {
		t.Errorf("Ожидаемый текст: %s, но получен: %s", text, string(data))
	}

	// Удаление тестового файла после теста
	os.Remove(filename)
}
