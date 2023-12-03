package transfer

import (
	"Paytabs/proj/user"
	"testing"
)

func TestNewTransfer(t *testing.T) {
	t.Log("Testing New Creating New Transfer")
	amount := 20.0
	source := user.User{Id: "1", Name: "1", Balance: 10.0}
	target := user.User{Id: "2", Name: "2", Balance: 20.0}

	new_transfer := NewTransfer(&source, &target, amount)

	if new_transfer.source != &source {
		t.Errorf("Creating new trannsfer failed")
	}
}

func TestParseInputAndReturnTransfer(t *testing.T) {
	testing_userMap := make(map[string]int)
	var testing_users []user.User

	source := user.User{Id: "1", Name: "1", Balance: 10.0}
	target := user.User{Id: "2", Name: "2", Balance: 20.0}

	testing_userMap["1"] = 0
	testing_userMap["2"] = 1

	testing_users = append(testing_users, source, target)

	want := NewTransfer(&source, &target, 20.0)
	need := ParseInputAndReturnTransfer("1", "2", "10.0", &testing_userMap, &testing_users)

	if *(want).source != source && *(need).source != source {
		t.Errorf("Creating new trannsfer from string inputs failed")
	}
}
