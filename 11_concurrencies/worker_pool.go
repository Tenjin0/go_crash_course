package main

import (
	"fmt"
	"math/rand"
)

func ExecuteWorkerPool() {

	type Job struct {
		id       int
		randomno int
	}

	type Result struct {
		job         Job
		sumofdigits int
	}

	jobs := make(chan Job, 10)
	// done := make(chan bool)
	// digits := func(number int) int {
	// 	sum := 0
	// 	no := number
	// 	for no != 0 {
	// 		digit := no % 10
	// 		sum += digit
	// 		no /= 10
	// 	}
	// 	time.Sleep(2 * time.Second)
	// 	return sum
	// }

	// worker := func(wg *sync.WaitGroup) {
	// 	wg.Done()
	// }

	allocate := func(noOfJobs int) {

		for i := 0; i < noOfJobs; i++ {
			randomno := rand.Intn(999)
			job := Job{id: i, randomno: randomno}
			jobs <- job
			fmt.Printf("new job sent %d with number %d\n", job.id, job.randomno)
		}
		close(jobs)
	}
	printJob := func() {

		for {
			job, ok := <-jobs
			if ok == false {
				break
			}
			fmt.Printf("Job id %d, input random %d, len: %d\n", job.id, job.randomno, len(jobs))
		}
	}
	go allocate(15)
	printJob()
	fmt.Println("Finished")
}
