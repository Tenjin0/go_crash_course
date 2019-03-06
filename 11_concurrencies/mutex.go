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
