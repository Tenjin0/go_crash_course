package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func tryToOpen() {
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println(err.Op, ": File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

func tryToLookUpHost() {

	addr, err := net.LookupHost("golanbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("Temporary error")
		} else {
			fmt.Println("generic error", err)
		}
	}
	fmt.Println(addr)
}

func tryToGlob() {
	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}

func main() {

	tryToOpen()
	tryToLookUpHost()
	tryToGlob()
}
