package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func generateNumbers(n int, done <-chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i <= n; i++ {
			select {
			case <-done:
				fmt.Println("Done channel completed")
				return
			default:
				out <- rand.Intn(n + 1)
			}
		}
	}()

	return out

}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// jab mutiple worker hotte hai toh pass as params channel for result
func singleWorker(input <-chan int, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range input {
		if isPrime(num) {
			result <- fmt.Sprintf("this number %d is prime", num)
		} else {
			result <- fmt.Sprintf("this number %d is not prime", num)
		}

	}

}

// send ka kaam kahtam hote hi close (unbuffered channel)
func FanOut(input <-chan int, result chan<- string) {
	defer close(result) //<----------
	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go singleWorker(input, result, &wg)
	}
	wg.Wait()
}

func fanIN(result chan string) {
	for value := range result {
		fmt.Println(value)
	}
}

func main() {
	done := make(chan int)
	defer close(done)

	numbersChan := generateNumbers(100, done)
	resultChan := make(chan string)
	go FanOut(numbersChan, resultChan)
	fanIN(resultChan)

}
