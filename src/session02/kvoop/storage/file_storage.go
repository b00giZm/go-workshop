package storage

import (
	"os"
	"bufio"
	"strings"
)

type FileStorage struct {
	path string
	mapper KeyValueMap
}

func NewFileStorage(name string) (*FileStorage, error) {
	fileStorage := &FileStorage{
		path: os.TempDir() + name,
		mapper: KeyValueMap{},
	}

	mapper, err := fileStorage.read()
	if err != nil {
		return nil, err
	}

	fileStorage.mapper = mapper

	return fileStorage, nil
}

func (f *FileStorage) Get(key string) (interface{}, bool) {
	value, found := f.mapper[key]
	return value, found
}

func (f *FileStorage) Set(key string, value string) {
	f.mapper[key] = value
	f.write(f.mapper)
}

func (f *FileStorage) All() KeyValueMap {
	return f.mapper
}

func (f *FileStorage) read() (KeyValueMap, error) {
	file, err := f.createOrOpen()
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	store := KeyValueMap{}
	if fileInfo.Size() <= 0 {
		return store, nil
	}

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

func (f* FileStorage) write(store KeyValueMap) error {
	file, err := f.createOrOpen()
	if err != nil {
		return err
	}

	defer file.Close()

	for key := range store {
		if _, err := file.WriteString(key + "=" + store[key] + "\n"); err != nil {
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
