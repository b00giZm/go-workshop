package storage

import (
	"os"
	"bufio"
	"strings"
)

type FileStorage struct {
	path string
	keyValueMap map[string]string
}

func NewFileStorage(path string) *FileStorage {
	storage := &FileStorage{
		path: os.TempDir() + path,
		keyValueMap: make(map[string]string),
	}
	storage.Read()

	return storage
}

func (f *FileStorage) Read() (map[string]string, error) {
	file, err := f.createOrOpen()
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	if fileInfo.Size() <= 0 {
		return f.keyValueMap, nil
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			parts := strings.Split(line, "=")
			f.keyValueMap[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return f.keyValueMap, nil
}

func (f *FileStorage) Write(key, value string) error {
	file, err := f.createOrOpen()
	if err != nil {
		return err
	}

	defer file.Close()

	f.keyValueMap[key] = value

	for key := range f.keyValueMap {
		if _, err := file.WriteString(key + "=" + f.keyValueMap[key] + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (f *FileStorage) createOrOpen() (*os.File, error) {
	if _, err := os.Stat(f.path); os.IsNotExist(err) {
		file, err := os.Create(f.path)
		if err != nil {
			return nil, err
		}

		return file, nil
	}

	file, err := os.OpenFile(f.path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}