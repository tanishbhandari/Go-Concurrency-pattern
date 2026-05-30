package main

import "fmt"

// send data from mutiple channels and combine into one recieve

func BrigeChannel(channels ...chan int) chan int {

	out := make(chan int)
	go func() {
		defer close(out)
		for _, ch := range channels {
			for val := range ch {
				out <- val
			}
		}
	}()

	return out
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	out := BrigeChannel(ch1, ch2, ch3)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	go func() {
		defer close(ch3)
		ch3 <- 5
		ch3 <- 6
	}()

	for msg := range out {
		fmt.Println(msg)
	}

}
