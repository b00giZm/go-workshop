package storage

import "fmt"

type KeyValueMap map[string]string

func (m KeyValueMap) String() string {
	var output string
	for key := range m {
		output += m.Print(key)
	}

	return output
}

func (m KeyValueMap) Print(key string) string {
	value, found := m[key]
	if !found {
		return ""
	}

	return fmt.Sprintf("> %s=%s\n", key, value)
}
