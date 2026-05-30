package main

import "fmt"

func activateGiftCards() func(int) int {

	amount := 100

	return func(debitAmount int) int {
		fmt.Println("remaing amount is ", amount, " now deducting ", debitAmount)
		amount -= debitAmount
		return amount
	}

}

func main() {
	giftCards1 := activateGiftCards()
	fmt.Println("giftCards1 rem amount is ", giftCards1(10))
	fmt.Println("giftCards1 rem amount is ", giftCards1(10))

	giftCards2 := activateGiftCards()
	fmt.Println("giftCards2 rem amount is ", giftCards2(10))
	fmt.Println("giftCards2 rem amount is ", giftCards2(10))
}
