package main

import (
	"fmt"
	"sync"
)

func CritalSectionWithoutMutex() {

	var x = 0
	increment := func(wg *sync.WaitGroup) {
		x = x + 1
		wg.Done()
	}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final value", x)

}

func CritalSectionWithMutex() {

	var x = 0
	increment := func(wg *sync.WaitGroup, m *sync.Mutex) {
		m.Lock()
		x = x + 1
		m.Unlock()
		wg.Done()
	}

	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &m)
	}

	wg.Wait()
	fmt.Println("Final value", x)

}

func CriticalSectionWithChan() {

	var x = 0

	increment := func(wg *sync.WaitGroup, ch chan bool) {

		ch <- true
		x++
		<-ch
		wg.Done()
	}

	var wg sync.WaitGroup
	var ch = make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, ch)
	}

	wg.Wait()
	fmt.Println("Final value", x)

}
