package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ExecuteWorkerPool(jobsToDo int, workersToExecuteJob int) {

	type Job struct {
		id       int
		randomno int
	}

	type Result struct {
		job         Job
		sumofdigits int
	}

	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	// done := make(chan bool)
	digits := func(number int) int {
		sum := 0
		no := number
		for no != 0 {
			digit := no % 10
			sum += digit
			no /= 10
		}
		time.Sleep(2 * time.Second)
		return sum
	}

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

	worker := func(i int, wg *sync.WaitGroup) {

		for job := range jobs {
			output := Result{job: job, sumofdigits: digits(job.randomno)}
			fmt.Printf("Calculate sum %d for job %d by worker id %d\n", output.sumofdigits, output.job.id, i)
			results <- output
		}
		wg.Done()
	}

	createWorkerPool := func(noOfWorkers int) {
		var wg sync.WaitGroup
		for i := 0; i < noOfWorkers; i++ {
			wg.Add(1)
			go worker(i, &wg)
		}
		wg.Wait()
		close(results)
	}

	result := func(done chan bool) {
		for result := range results {
			fmt.Printf("Job is %d, input random no %d, sum %d READ RESULT\n", result.job.id, result.job.randomno, result.sumofdigits)
		}

		done <- true
	}

	startTime := time.Now()

	allocate(jobsToDo)

	done := make(chan bool)
	go result(done)
	createWorkerPool(workersToExecuteJob)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
