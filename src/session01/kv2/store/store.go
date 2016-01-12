package store

import (
	"os"
	"bufio"
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

	store := Store{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			parts := strings.Split(line, "=")
			store[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
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