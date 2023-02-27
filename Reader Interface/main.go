package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	data, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, data)
}
