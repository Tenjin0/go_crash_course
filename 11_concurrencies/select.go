package main

import (
	"fmt"
	"time"
)

func Select() {

	server1 := func(ch chan string) {
		time.Sleep(6 * time.Second)
		ch <- "from server 1"
	}

	server2 := func(ch chan string) {
		time.Sleep(3 * time.Second)
		ch <- "from server 2"
	}

	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)

loop:
	for {
		time.Sleep(1 * time.Second)
		select {
		case s1 := <-output1:
			fmt.Println(s1)
			break loop
		case s2 := <-output2:
			fmt.Println(s2)
			break loop
		default:
			fmt.Println("Nothing to be done")
		}
	}
	fmt.Println("The End")
}
