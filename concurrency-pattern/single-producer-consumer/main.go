package main

import (
	"fmt"
	"sync"
)

// Producer function to increment the value and send it to the channel
func increment(prev int, data chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	data <- prev + 1
}

// Consumer function to read values from the channel
func consumer(data chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range data {
		fmt.Println("val is", val)
	}
}

func main() {
	data := make(chan int)
	var producerWg sync.WaitGroup

	// Launch producers
	for i := 0; i < 10; i++ {
		producerWg.Add(1)
		go increment(i, data, &producerWg)
	}

	// Close the data channel when all producers are done
	go func() {
		producerWg.Wait()
		close(data)
	}()

	var consumerWg sync.WaitGroup

	// Launch consumers
	for i := 0; i < 10; i++ {
		consumerWg.Add(1)
		go consumer(data, &consumerWg)
	}

	// Wait for all consumers to finish
	consumerWg.Wait()
}
