package main

import (
	"fmt"
	"sync"
)

type AccountBalance struct {
	balance int
	mu      sync.Mutex
}

func (ab *AccountBalance) Deposit(amount int) {
	ab.mu.Lock()
	defer ab.mu.Unlock()
	ab.balance += amount

	fmt.Printf("Deposit %d done\n", amount)
}

func (ab *AccountBalance) Withdraw(amount int) {
	ab.mu.Lock()
	defer ab.mu.Unlock()

	if ab.balance < amount {
		fmt.Println("Insufficent balance")
	}

	ab.balance -= amount
	fmt.Println("Withdrawal done")
}

func (ab *AccountBalance) GetBalance() int {
	ab.mu.Lock()
	defer ab.mu.Unlock()
	return ab.balance
}

func main() {

	/*
When is Mutex Useful in Go?
SituationNeed Mutex?
Shared variable read/written by multiple goroutines✅ Yes
Simple sequential code, no goroutines❌ No
Using Go channels to pass data❌ Usually no
Read-heavy workload🔄 Use sync.RWMutex instead


The Go mantra on this:

"Don't communicate by sharing memory; share memory by communicating" — meaning Go channels are often preferred over mutexes for passing data between goroutines, but mutexes are the right tool when you have a shared resource like this

*/

	// counter := 0

	var wg sync.WaitGroup

	var account = AccountBalance{balance: 100}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(amount int) {
			account.Deposit(amount)
			defer wg.Done()

		}(i + 1)

	}

	wg.Wait()

	fmt.Println(account.GetBalance())


	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)

	// 	go func() {
	// 		mutex.Lock()
	// 		defer mutex.Unlock()
	// 		counter++
	// 		println(counter)
	// 		defer wg.Done()
	// 	}()

	// }

	// wg.Wait()

	//fmt.Println(counter);

}


