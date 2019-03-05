package main

import (
	"fmt"
	"time"
)

func BuffereChannel() {

	ch := make(chan string, 2)
	fmt.Println("Length is", len(ch))
	ch <- "naveen"
	fmt.Println("Length is", len(ch))
	ch <- "paul"
	fmt.Println("Length is", len(ch))
	// ch <- "dave"
	fmt.Println("Capacity is", cap(ch))
	fmt.Println(<-ch)
	fmt.Println("Are we blocked ? not yet.", "Length is", len(ch))
	fmt.Println(<-ch)
	fmt.Println("Are we blocked ? not yet.", "Length is", len(ch))

	// ch2 := make(chan int, 2)
	// go write(ch2)

	// for v := range ch2 {
	// 	fmt.Println("read value", v, "from ch")
	// 	// time.Sleep(2 * time.Second)

	// }
}

func write(ch chan<- int) {

	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("SuccessFully wrote", i, "to ch")
		time.Sleep(2 * time.Second)
	}
	close(ch)
}
