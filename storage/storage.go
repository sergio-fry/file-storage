package storage

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

type Storage struct {
	path string
}

func (s Storage) write(key string, content io.Reader) {
	file, err := os.Create(s.fileNameFromKey(key))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.ReadFrom(content)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Storage) read(key string) io.ReadCloser {
	file, err := os.Open(s.fileNameFromKey(key))

	if err != nil {
		log.Fatal(err)
	}
	return file
}

func (s Storage) fileNameFromKey(key string) string {
	hash := md5.Sum([]byte(key))
	keymd5 := hex.EncodeToString(hash[:])

	return s.path + "/" + keymd5
}
