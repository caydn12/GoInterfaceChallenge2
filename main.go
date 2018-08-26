package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileLocation := os.Args[1]

	file, err := os.Open(fileLocation)

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(-1)
	}

	io.Copy(os.Stdout, file)
}
