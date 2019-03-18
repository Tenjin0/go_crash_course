package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	fmt.Println("test init")
}
func readFromAbsolutePath() {
	data, err := ioutil.ReadFile("/home/tenji/go/src/github.com/Tenjin0/go_crash_course/19_reading_files/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
func readRelativePath() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func readPathFromCommandLine() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func readChunkFiles() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	close := func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}

	defer close()

	r := bufio.NewReader(f)
	b := make([]byte, 3)

	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file", err)
			break
		}
		fmt.Println(string(b))
	}
}

func readLineByLine() {

	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	// box := packr.NewBox("../19_reading_files")
	// data := box.String("test.txt")
	// fmt.Println("Contents of file", data)
	// readChunkFiles()
	readLineByLine()
	// readPathFromCommandLine()
}
