package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ExecuteTask(goRoutineId int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("goRountine %d is executing tasks number %d\n", goRoutineId, task)
		fmt.Printf("goRountine %d is Completed tasks number %d\n", goRoutineId, task)
	}

}

func Example1() {
	// 3 go routine and 10 tasks we have

	x := time.Now()
	taskNumber := 10
	goRoutineNumber := 3

	tasks := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= goRoutineNumber; i++ {
		wg.Add(1)
		go ExecuteTask(i, tasks, &wg)
	}

	for i := 1; i <= taskNumber; i++ {
		tasks <- i
	}
	close(tasks)

	wg.Wait()
	fmt.Println("time taken is ", time.Since(x))
}

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job              Job
	sumofdigits, wId int
}

func allocateJobs() chan Job {
	// Create 50 jobs. Assign each one a job number (id)
	//and a random integer (randomno) using Job struct,
	//and pass them to the jobs channel.
	jobsChan := make(chan Job)
	go func() {
		for i := 1; i <= 50; i++ {
			randNo := rand.Intn(100) + 1
			job := Job{
				id:       i,
				randomno: randNo,
			}
			jobsChan <- job
		}
		close(jobsChan)
	}()

	return jobsChan
}

func WorkerChannel(jobsChan chan Job) chan Result {
	// Create a pool of 10 workers.
	resultChan := make(chan Result)
	go func() {
		var wg sync.WaitGroup
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go worker(i, jobsChan, resultChan, &wg)
		}
		wg.Wait()

		close(resultChan)
	}()

	return resultChan

}
func Example2() {

	// Create a buffer channel of size 10.
	jobsChan := allocateJobs()
	resultChan := WorkerChannel(jobsChan)

	for result := range resultChan {
		fmt.Println(fmt.Sprintf("{jobId :%d , randomNumber: %d} and {sum : %d , workerID: %d}", result.job.id, result.job.randomno, result.sumofdigits, result.wId))
	}

}

// Create a function that calculates the sum of the interger’s digits passed as a parameter.
func calculateSum(job Job) int {
	sum := 0
	num := job.randomno
	for num > 0 {
		rem := num % 10
		sum += rem
		num = num / 10

	}
	return sum
}

// Create a function to store the values in struct Result
// and pass them into the results channel.
// This work should be done by a pool of workers.
func worker(workerId int, jobsChan chan Job, resultChan chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobsChan {
		sum := calculateSum(job)

		resultChan <- Result{
			job:         job,
			sumofdigits: sum,
			wId:         workerId,
		}
	}
}

func main() {
	Example2()
}
