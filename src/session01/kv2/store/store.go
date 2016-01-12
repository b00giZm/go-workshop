package store

import (
	"os"
	"bufio"
	"strings"
)

type Store map[string]string

func Open(name string) (Store, error) {
	file, err := createOrOpen(os.TempDir() + name)
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	if fileInfo.Size() <= 0 {
		return Store{}, nil
	}

	defer file.Close()

	if fileInfo.Size() <= 0 {
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

func Write(name string, store Store) (bool, error) {
	file, err := createOrOpen(os.TempDir() + name)
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

func createOrOpen(name string) (*os.File, error) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			return nil, err
		}

		return file, nil
	}

	file, err := os.OpenFile(name, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}