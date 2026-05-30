package main

import (
	"fmt"
	"math/rand"
)

// Write a program that generates n random integers within the range of n. For example, if n is 5, generate five random integers less than or equal to 5 by following these steps:

// Pass n as a parameter.
// Return a type struct holding integer n, and string it on a new line with the message Your random number is x.
// Print the log.
// Write a consumer function which prints n times or can use the range function.

type NumberMessage struct {
	number  int
	message string
}

func generateNumber(n int, out chan<- NumberMessage, done <-chan int) {
	defer close(out)
	for i := 0; i <= n; i++ {
		select {
		case <-done:
			fmt.Println("Done channel completed")
			return
		default:
			x := rand.Intn(n + 1)

			out <- NumberMessage{
				number:  x,
				message: fmt.Sprintf("Current Number is %d", x),
			}
		}
	}

}
func handleGeneration(n int, done <-chan int) chan NumberMessage {
	out := make(chan NumberMessage)
	go generateNumber(n, out, done)

	return out
}

func main() {
	// done := make(chan int)
	// defer close(done)
	// for num := range handleGeneration(5, done) {
	// 	fmt.Println("Number : ", num)
	// 	fmt.Println("Message : ", num)
	// }

	Generics()
}
