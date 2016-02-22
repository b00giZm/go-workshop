package main

import (
	"os"
	"strings"
)

func main() {
	store := NewDefaultStore()

	defer store.Close()

	isReadMode := true
	readOps := make([]string, 0)
	for _, arg := range os.Args[1:] {
		parts := strings.Split(arg, "=");
		if len(parts) <= 1 {
			// Read operation
			readOps = append(readOps, parts[0])
		} else {
			// Write operation
			isReadMode = false
			store.Add(parts[0], parts[1])
		}
	}

	if isReadMode {
		len := len(readOps)
		if (len < 1) {
			store.All()
		} else {
			for _, key := range readOps {
				store.Get(key)
			}
		}
	}
}
