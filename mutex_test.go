package go_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// to avoid race condition
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter", x)
}

type BankAccount struct {
	RWMutext sync.RWMutex
	Balance  int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutext.Lock()
	account.Balance += amount
	account.RWMutext.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutext.RLock()
	balance := account.Balance
	account.RWMutext.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Balance", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println(user1.Name, "user1 Locked")
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println(user2.Name, "user2 Locked")
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "User 1",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "User 2",
		Balance: 2000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user1, &user2, 200000)

	time.Sleep(3 * time.Second)

	fmt.Println("User 1", user1.Name, "Balance", user1.Balance)
	fmt.Println("User 2", user2.Name, "Balance", user2.Balance)
}
