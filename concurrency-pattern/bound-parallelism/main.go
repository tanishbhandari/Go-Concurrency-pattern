package main

import (
	"fmt"
	"sync"
	"time"
)

// use buffered sem for limit worker
func worker(task int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Executing task ", task)
	time.Sleep(3 * time.Millisecond)
}
func BoundedParallelism(tasks []int, workerNo int) {
	sem := make(chan struct{}, workerNo)
	var wg sync.WaitGroup

	for _, task := range tasks {
		sem <- struct{}{}
		wg.Add(1)
		go func(t int) {
			defer func() { <-sem }()
			worker(t, &wg)
		}(task)
	}

	wg.Wait()
	close(sem)
}
func main() {
	task := []int{1, 2, 3, 4, 5, 6, 7}
	BoundedParallelism(task, 3)
}
