package main

import (
	"errors"
	"log"
	"os"
)

type FileNotFoundError error
type InternalServerError error

var (
	ErrFileNotFound FileNotFoundError = errors.New("file not found")
	ErrInternalServerError InternalServerError = errors.New("internal server error")
)

func GetFileContent(path string) ([]byte, error) {
	actualPath := filesDir + "/" + path
	if _, err := os.Stat(actualPath); errors.Is(err, os.ErrNotExist) {
		return nil, ErrFileNotFound
	}

	body, err := os.ReadFile(actualPath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		return nil, ErrInternalServerError
	}

	return body, nil
}

func writeFile(filename string, data []byte, contentLength int) error {
	actualPath := filesDir + "/" + filename

	err := os.WriteFile(actualPath, data[:contentLength], 0644)
    if err != nil {
        return ErrInternalServerError
    }
    return nil
}
