package main

import (
	"fmt"
	"math/rand"
)

func generateIntegers(n int) <-chan int {

	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			num := rand.Intn(n) + 1
			out <- num
		}
		close(out)
	}()
	return out
}

func squareGeneratedIntegers(input <-chan int) <-chan int {

	output := make(chan int)
	go func() {
		for val := range input {
			output <- val
		}
		close(output)
	}()
	return output
}

// generate integers
// square generated interger
// display all squared integers
func main() {
	numbersChan := generateIntegers(5)

	outputChan := squareGeneratedIntegers(numbersChan)

	for result := range outputChan {
		fmt.Println("result no is ", result)
	}

}
