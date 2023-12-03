package user

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	user := User{Id: "1", Name: "1", Balance: 10.0}

	user.Withdraw(1.0)

	if user.Balance != 9.0 {
		t.Errorf("Withdraw test failed")
	}
}
