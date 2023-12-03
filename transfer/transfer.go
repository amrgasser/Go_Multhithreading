package transfer

import (
	"Paytabs/proj/user"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Transfer struct {
	source *user.User
	target *user.User
	amount float64
}

var transfers []Transfer

func NewTransfer(source *user.User, target *user.User, amount float64) *Transfer {
	transfer := Transfer{source: source, target: target, amount: amount}
	return &transfer
}

func ParseInputAndReturnTransfer(sourceId string, targetId string, amountString string, userMap *map[string]int, users *[]user.User) *Transfer {

	sourceInd, exists := (*userMap)[sourceId]
	if !exists {
		fmt.Println("Source Id is incorrect")
		return nil
	}
	targetInd, exists := (*userMap)[targetId]
	if !exists {
		fmt.Println("Target Id is incorrect")
		return nil
	}
	_source := &(*users)[sourceInd]
	_target := &(*users)[targetInd]

	_amount, err := strconv.ParseFloat(amountString, 64)

	if err != nil {
		fmt.Println("Invalid amount!")
		return nil
	}

	return NewTransfer(_source, _target, _amount)
}

func PerformTransfer(t *Transfer, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	if t.source.Withdraw(t.amount) {
		t.target.Deposit(t.amount)
		ch <- fmt.Sprintf("from %s to %s with amount: %f Succeeded at %d\n", t.source.Name, t.target.Id, t.amount, time.Now().UnixNano())
	} else {
		ch <- fmt.Sprintf("from %s to %s with amount: %f Failed at %d\n", t.source.Name, t.target.Id, t.amount, time.Now().UnixNano())
	}
	go PrintFromChannel(ch, wg)
}

func PrintFromChannel(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-ch)
}
