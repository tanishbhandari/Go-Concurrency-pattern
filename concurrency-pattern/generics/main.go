package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func sumNumbers[T Number](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	nums1 := []int32{1, 2, 3, 4, 5}
	nums2 := []float32{1.1, 2.34, 32.4, 3}

	fmt.Println(sumNumbers(nums1))
	fmt.Println(sumNumbers(nums2))

}
