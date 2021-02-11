package main

import (
	"fmt"
	"github.com/netwar1994/channels/pkg/card"
)

func main() {
	transaction1 := card.Transaction{
		UserId: 0,
		Sum:    1000,
		MCC:    "5411",
	}
	transaction2 := card.Transaction{
		UserId: 2,
		Sum:    2000,
		MCC:    "5812",
	}
	transaction3 := card.Transaction{
		UserId: 2,
		Sum:    10,
		MCC:    "5816",
	}
	transaction4 := card.Transaction{
		UserId: 2,
		Sum:    100,
		MCC:    "5816",
	}

	transactions1 := []card.Transaction{transaction1, transaction2, transaction2, transaction3, transaction3, transaction3, transaction4}

	fmt.Println("Функция без горутин")
	fmt.Println("----------------------------------")
	for k, v := range card.SumByCategory(transactions1) {
		fmt.Println(k, v)
	}
	fmt.Println("\nФункция с mutex'ом")
	fmt.Println("----------------------------------")
	for k, v := range card.SumByCategoryMutex(transactions1, 10) {
		fmt.Println(k, v)
	}
	fmt.Println("\nФункция с каналами")
	fmt.Println("----------------------------------")
	for k, v := range card.SumByCategoryChannels(transactions1, 10) {
		fmt.Println(k, v)
	}
	fmt.Println("\nФункция с mutex'ом без вызова фунции 1")
	fmt.Println("----------------------------------")
	for k, v := range card.SumByCategoryMutexWithoutFunc(transactions1, 10) {
		fmt.Println(k, v)
	}

}
