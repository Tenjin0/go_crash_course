package main

import (
	"fmt"
	"os"
)

func WriteStringInFile() {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeByteInFile() {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeLineInFile() {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"Welcome", "golang", "the end"}

	for _, v := range d {

		_, err := fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")

}

func appendToFile() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "File handling is easy"
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file  appended successfully")
}

func writeConcurrently() {

}

func main() {
	writeConcurrently()
}
