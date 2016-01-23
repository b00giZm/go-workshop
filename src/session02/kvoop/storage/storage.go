package storage

type Storage interface {
	Get(key string) (interface{}, bool)
	Set(key string, value string)
	All() KeyValueMap
}
