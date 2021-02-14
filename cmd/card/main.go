package main

import (
	"fmt"
	"github.com/netwar1994/channels/pkg/card"
	"log"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	transactions := card.MakeTransactions(1)
	fmt.Println("Функция без горутин")
	fmt.Println("----------------------------------")
	for k, v := range card.SumByCategory(transactions, 1) {
		fmt.Println(k, v)
	}
	fmt.Println("\nФункция с mutex'ом")
	fmt.Println("----------------------------------")
	for k, v := range card.SumByCategoryMutex(transactions, 1, 10) {
		fmt.Println(k, v)
	}
	fmt.Println("\nФункция с каналами")
	fmt.Println("----------------------------------")
	for k, v := range card.SumByCategoryChannels(transactions, 1, 10) {
		fmt.Println(k, v)
	}
	fmt.Println("\nФункция с mutex'ом без вызова фунции 1")
	fmt.Println("----------------------------------")
	for k, v := range card.SumByCategoryMutexWithoutFunc(transactions, 1, 10) {
		fmt.Println(k, v)
	}

}
