package card

import (
	"sync"
)

type Transaction struct {
	UserId int64
	Sum    int64
	MCC    string
}

func SumByCategory(transactions []Transaction) map[string]int64 {
	result := make(map[string]int64)

	for _, transaction := range transactions {
		getMMC := TranslateMCC(transaction.MCC)
		result[getMMC] += transaction.Sum
	}
	return result
}

func SumByCategoryMutex(transactions []Transaction, goroutines int) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)
	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			m := SumByCategory(part)

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

func SumByCategoryChannels(transactions []Transaction, goroutines int) map[string]int64 {
	result := make(map[string]int64)
	ch := make(chan map[string]int64)
	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func(ch chan<- map[string]int64) {
			ch <- SumByCategory(part)
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

func SumByCategoryMutexWithoutFunc(transactions []Transaction, goroutines int) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)
	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			for _, t := range part {
				mu.Lock()
				result[TranslateMCC(t.MCC)] += t.Sum
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return result
}

