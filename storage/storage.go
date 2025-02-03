package storage

import (
	"log"
	"os"
)

func writeToFile(key string, text string) {
	file, err := os.Create(fileNameFromKey(key))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal(err)
	}
}

func fileNameFromKey(key string) string {
	return key
}
