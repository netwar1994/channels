package main

import (
	"github.com/netwar1994/channels/pkg/card"
	"fmt"
	"time"
)

func main()  {
	transaction1 := card.Transaction{
		Id:     0,
		Sum:    1000,
		Date:   time.Now().Format(time.Stamp),
		MCC:    "5411",
		Status: "pending",
	}
	transaction2 := card.Transaction{
		Id:     1,
		Sum:    2000,
		Date:   time.Now().Format(time.Stamp),
		MCC:    "5812",
		Status: "pending",
	}
	transaction3 := card.Transaction{
		Id:     2,
		Sum:    10,
		Date:   time.Now().Format(time.Stamp),
		MCC:    "5816",
		Status: "pending",
	}

	transactions1 := []card.Transaction{transaction1, transaction2, transaction3, transaction3, transaction3 }
	//transactions2 := []card.Transaction{transaction2, transaction3,transaction1, transaction2, transaction3,transaction1, transaction2, transaction3,transaction1, transaction3, transaction1, transaction3}

	//card1 := card.Card{
	//	Id:           1,
	//	Owner: 	"User1",
	//	Transactions: transactions1,
	//}
	//
	//card2 := card.Card{
	//	Id:           2,
	//	Owner: 	"User2",
	//	Transactions: transactions2,
	//}

//	transactions := []card.Transaction{}
//	cards := []card.Card{card1, card2}

	for k, v := range card.ExpensesByCategory(transactions1) {
		fmt.Println(k, v)
	}

	for k, v := range card.ExpensesByCategoryMutex(transactions1) {
		fmt.Println(k, v)
	}

}