package user

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Id      string
	Name    string
	mu      sync.Mutex
	Balance float64
}

func (user *User) Withdraw(amount float64) bool {
	time.Sleep(1000 * time.Millisecond)
	user.mu.Lock()
	if user.Balance >= amount {
		user.Balance -= amount
		user.mu.Unlock()
		return true
	}
	user.mu.Unlock()
	return false
}

func (user *User) Deposit(amount float64) {
	user.mu.Lock()
	user.Balance += amount
	user.mu.Unlock()
}

// Load Data From json to array and convert Balance as a string to a float
func (user *User) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Id      string
		Name    string
		Balance string
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	user.Name = tmp.Name
	user.Id = tmp.Id
	// Convert the string to float
	Balance, err := strconv.ParseFloat(tmp.Balance, 64)
	if err != nil {
		return err
	}
	user.Balance = Balance
	user.mu = sync.Mutex{}

	return nil
}
