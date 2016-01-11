package store

import (
	"os"
	"strings"
	//"fmt"
)

type Store map[string]string

func Open(filename string) (Store, error) {
	storePath := os.TempDir() + filename
	fileInfo, err := ensureFile(storePath);
	if err != nil {
		return nil, err
	}

	file, err := os.Open(storePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileSize := fileInfo.Size()
	if fileSize <= 0 {
		return Store{}, nil
	}

	// Reading the whole db into a byte array. Do not do this in a real world application!
	bytes := make([]byte, fileSize)
	if _, err := file.Read(bytes); err != nil {
		return nil, err
	}

	store := Store{}
	for _, line := range strings.Split(string(bytes), "\n") {
		if len(line) > 0 {
			parts := strings.Split(line, "=")
			store[parts[0]] = parts[1]
		}
	}

	return store, nil
}

func Write(filename string, store Store) (bool, error) {
	storePath := os.TempDir() + filename
	file, err := os.Create(storePath)
	if err != nil {
		return false, err
	}

	defer file.Close()

	for key := range store {
		if _, err := file.WriteString(key + "=" + store[key] + "\n"); err != nil {
			return false, err
		}
	}

	return true, nil
}

func ensureFile(filePath string) (os.FileInfo, error) {
	fileInfo, err := os.Stat(filePath);
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}

		fileInfo, err = file.Stat();
		if err != nil {
			return nil, err
		}

		return fileInfo, nil
	}

	return fileInfo, nil
}