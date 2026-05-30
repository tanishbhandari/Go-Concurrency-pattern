package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {

	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing Work ", i)
			i++
		}
	}
}
func main() {
	goChannel := make(chan int)
	for {
		select {
		case res := <-goChannel:
			fmt.Println(res)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
		}
	}

}
