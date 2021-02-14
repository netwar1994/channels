package card

import (
	"sync"
)

type Transaction struct {
	UserId int64
	Sum    int64
	MCC    string
}

func MakeTransactions(userId int64) []Transaction {
	const usersCount = 1_000
	const transactionsCount = 1_000
	const transactionAmount = 1_00
	transactions := make([]Transaction, usersCount*transactionsCount)
	x := Transaction{
		UserId: userId,
		Sum:    transactionAmount,
		MCC:    "5411",
	}
	y := Transaction{
		UserId: userId,
		Sum:    transactionAmount,
		MCC:    "5812",
	}
	z := Transaction{
		UserId: 2,
		Sum:    transactionAmount,
		MCC:    "5812",
	}

	for index := range transactions {
		switch index % 100 {
		case 0:
			transactions[index] = x
		case 20:
			transactions[index] = y
		default:
			transactions[index] = z
		}
	}
	return transactions
}

func SumByCategory(transactions []Transaction, userId int64) map[string]int64 {
	result := make(map[string]int64)

	for _, transaction := range transactions {
		if transaction.UserId == userId {
			getMMC := TranslateMCC(transaction.MCC)
			result[getMMC] += transaction.Sum
		}
	}
	return result
}

func SumByCategoryMutex(transactions []Transaction, userId int64, goroutines int) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)
	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			m := SumByCategory(part, userId)

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

func SumByCategoryChannels(transactions []Transaction, userId int64, goroutines int) map[string]int64 {
	result := make(map[string]int64)
	ch := make(chan map[string]int64)
	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func(ch chan<- map[string]int64) {
			ch <- SumByCategory(part, userId)
		}(ch)
	}
	finished := 0
	for value := range ch {
		for k, v := range value {
			result[k] += v
		}
		finished++
		if finished == goroutines {
			break
		}
	}
	return result
}

func SumByCategoryMutexWithoutFunc(transactions []Transaction, userId int64, goroutines int) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)
	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			for _, t := range part {
				if t.UserId == userId {
					mu.Lock()
					result[TranslateMCC(t.MCC)] += t.Sum
					mu.Unlock()
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return result
}
