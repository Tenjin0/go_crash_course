package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("Hello world  goroutine: 1")
	done <- true
	fmt.Println("Hello world  goroutine: 2")

}

func channel() {
	var a chan bool

	if a == nil {
		fmt.Println("channel a is nill, going to define it")
		a = make(chan bool)
		fmt.Printf("Type of a is %T\n", a)
	}

	go hello(a)
	fmt.Println("Channel function: 1")
	<-a // block the main go routine and force it to wait
	fmt.Println("Channel function: 2")
}
