package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateNumbers[T any, K any](done <-chan K, fn func() T) <-chan T {
	streamData := make(chan T)

	go func() {
		defer close(streamData)
		for {
			select {
			case <-done:
				return
			case streamData <- fn():
			}
		}
	}()

	return streamData
}

func Generics() {
	done := make(chan bool)
	randomIntegers := func() int { return rand.Intn(1000) }

	dataChan := GenerateNumbers(done, randomIntegers)

	go func() {
		defer close(done)
		time.Sleep(3 * time.Second)
	}()

	for data := range dataChan {
		fmt.Println("Data is ", data)
	}
}
