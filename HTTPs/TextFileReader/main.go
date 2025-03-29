package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// read input in the command line as argument
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// f has Read(b []byte) method!
	// so f is an interface of Reader!
	io.Copy(os.Stdout, f)

}
