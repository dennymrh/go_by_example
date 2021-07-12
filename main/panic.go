package main

import "os"

func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

// Running this program will cause it to panic,
// print an error message and goroutine traces,
// and exit with a non-zero status.
