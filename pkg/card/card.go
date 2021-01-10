package card

import (
	"sync"
)

type Transaction struct {
	Id int64
	Sum int64
	Date string
	MCC string
	Status string
}

type Card struct {
	Id int64
	Owner string
	Issuer string
	Balance int64
	Currency string
	Number string
	Transactions []Transaction
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Balance -= transaction.Sum
}

func ExpensesByCategory(transactions []Transaction) map[string]int64 {
	categories := make(map[string]int64)

	for _, transaction := range transactions {
		getMMC := TranslateMCC(transaction.MCC)
		categories[getMMC] += transaction.Sum
	}

	return categories
}

func ExpensesByCategoryMutex(transactions []Transaction) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)
	const partsCount = 2
	partSize := len(transactions) / partsCount

	for i:=0; i<partsCount; i++ {
		wg.Add(1)
	    part := transactions[i*partSize:(i+1)* partSize]
		go func() {
			m := ExpensesByCategory(part)

			mu.Lock()
			for k, v := range m {
				result[k] += v
			}

			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return result
}