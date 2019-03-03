package main

import (
	"fmt"
	"time"
)

func Buffered_channel() {

	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println("Are we bloked ? not yet")
	fmt.Println(<-ch)
	fmt.Println("Are we bloked ? not yet")

	ch2 := make(chan int, 2)
	go write(ch2)

	for v := range ch2 {
		fmt.Println("read value", v, "from ch")
		// time.Sleep(2 * time.Second)

	}
}

func write(ch chan<- int) {

	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("SuccessFully wrote", i, "to ch")
		time.Sleep(2 * time.Second)
	}
	close(ch)
}
