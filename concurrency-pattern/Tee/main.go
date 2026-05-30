package main

import (
	"fmt"
	"sync"
)

//single input and mutiple output channels with same message

func Tee(input chan string) (chan string, chan string) {

	out1 := make(chan string)
	out2 := make(chan string)
	go func() {
		defer close(out1)
		defer close(out2)
		for msg := range input {
			out1 <- msg
			out2 <- msg
		}
	}()

	return out1, out2
}
func main() {
	input := make(chan string)

	out1, out2 := Tee(input)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for out := range out1 {
			fmt.Println("output1 : ", out)
		}
	}()

	go func() {
		defer wg.Done()
		for out := range out2 {
			fmt.Println("output2 : ", out)
		}
	}()

	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("msg no is %d", i)
		input <- msg
	}

	close(input)
	wg.Wait()
}
