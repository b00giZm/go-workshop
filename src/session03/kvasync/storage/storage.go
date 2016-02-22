package storage

type Storage interface {
	Read() (map[string]string, error)
	Write(key, value string) error
}
