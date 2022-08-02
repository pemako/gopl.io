package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args[0:] {
		fmt.Printf("index=%d vlaue=%s\n", k, v)
	}
}
