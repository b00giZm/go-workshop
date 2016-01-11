package main

import (
	"os"
	"fmt"
	"strings"
	"session01/kv/store"
)

const storeFilename string = "store.db"

func main() {

	s, err := store.Open(storeFilename)
	if err != nil {
		panic(err)
	}

	writeMode := false
	toRead := make([]string, 0)
	for _, arg := range os.Args[1:len(os.Args)] {
		parts := strings.Split(arg, "=");
		if len(parts) <= 1 {
			toRead = append(toRead, arg)
		} else {
			writeMode = true
			s[parts[0]] = parts[1]
		}
	}

	if writeMode && len(toRead) <= 0 {
		if _, err = store.Write(storeFilename, s); err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	if len(toRead) > 0 {
		for _, key := range toRead {
			value, exists := s[key]
			if exists {
				fmt.Println("> " + key + "=" + value)
			}
		}

		os.Exit(0)
	}

	for key := range s {
		fmt.Println("> " + key + "=" + s[key])
	}
}
