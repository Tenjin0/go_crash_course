package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {

	process := func(x int, wg *sync.WaitGroup) {
		fmt.Println("Goroutine started", x)
		time.Sleep(2 * time.Second)
		fmt.Printf("Goroutine %d ended\n", x)
		wg.Done()
	}

	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
