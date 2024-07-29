package main

import (
	"errors"
	"log"
	"os"
)

type FileNotFoundError error

var (
	ErrFileNotFound FileNotFoundError = errors.New("file not found")
)

func GetFileContent(path string) ([]byte, error) {
	actualPath := filesDir + "/" + path
	if _, err := os.Stat(actualPath); errors.Is(err, os.ErrNotExist) {
		return nil, ErrFileNotFound
	}

	body, err := os.ReadFile(actualPath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return body, nil
}
