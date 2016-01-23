package storage

import (
	"fmt"
	"sort"
)

type KeyValueMap map[string]string

func (m KeyValueMap) String() string {
	var output string

	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
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
