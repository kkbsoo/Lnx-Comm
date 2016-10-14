package main

import (
	"os"
	"syscall"
)

func main() {
	sOpenFile := os.Args[1]
	sCloseFile := os.Args[2]

	err := os.Link(sOpenFile, sCloseFile)
	if err != nil {
		panic(err)
	}

	err = syscall.Unlink(sOpenFile)
	if err != nil {
		panic(err)
	}
}