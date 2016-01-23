package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	kv, err := NewDefaultStore()
	if err != nil {
		panic(err)
	}

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
			kv.Set(parts[0], parts[1])
		}
	}

	if isReadMode {
		len := len(readOps)
		if (len < 1) {
			fmt.Print(kv.All())
		} else {
			fmt.Print(kv.GetMultiple(readOps))
		}
	}
}
