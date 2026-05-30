package main

import (
	"fmt"
	"sync"
)

func ManageTickets(done chan struct{}, ticketReqChan chan int, totalTickets *int, result chan string) {
	defer close(result)

	for {
		select {
		case <-done:
			return
		default:
			userId := <-ticketReqChan
			if *totalTickets > 0 {
				*totalTickets -= 1
				//fmt.Println(fmt.Sprintf("user Id %d bought %d tickets", NoOfTickets, NoOfTickets))
				result <- fmt.Sprintf("user Id %d bought %d tickets and remaining %d", userId, 1, *totalTickets)

			} else {
				//fmt.Println(fmt.Sprintf("tickets not available for user id %d", NoOfTickets))
				result <- fmt.Sprintf("tickets not available for user id %d and remaining %d", userId, *totalTickets)
			}
		}
	}

}

func buyTickets(userID int, ticketReqChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ticketReqChan <- userID
}

// we have 500 ticket to sell
// but users can be 1000 +
//so want to run program concurrently shared resource issue

// solve-->use only one goroutine to write
// take requests from multiple goroutine
func main() {
	// done := make(chan struct{})
	// defer close(done)

	// totalTickets := 100
	// ticketReqChan := make(chan int) // taking requests

	// resultChan := make(chan string)
	// go func() {
	// 	for value := range resultChan {
	// 		fmt.Println(value)
	// 	}
	// }()

	// go ManageTickets(done, ticketReqChan, &totalTickets, resultChan)

	// var wg sync.WaitGroup
	// for userID := 1; userID <= 200; userID++ {
	// 	wg.Add(1)
	// 	go buyTickets(userID, ticketReqChan, &wg)
	// }
	// wg.Wait()
	// close(ticketReqChan)
	WithOutConfinement()

}

func buyTicketsWithoutConfinemnt(userID int, totalTickets *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done() // Mark this goroutine as done when it finishes

	mu.Lock() // Lock the critical section
	if *totalTickets > 0 {
		*totalTickets-- // Sell one ticket
		fmt.Printf("Ticket sold! Tickets remaining: %d\n", *totalTickets)
	}
	mu.Unlock() // Unlock the critical section

}
func WithOutConfinement() {
	totalTickets := 100

	var wg sync.WaitGroup
	var mu sync.Mutex
	for user := 1; user <= 200; user++ {
		wg.Add(1)
		go buyTicketsWithoutConfinemnt(user, &totalTickets, &wg, &mu)
	}
	wg.Wait()
}
