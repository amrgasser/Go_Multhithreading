package load

import (
	"Paytabs/proj/user"
	"testing"
)

func TestLoadData(t *testing.T) {
	testing_userMap := make(map[string]int)
	var testing_users []user.User

	LoadData("../test.json", &testing_userMap, &testing_users)
	_, exists := testing_userMap["1"]

	if len(testing_users) != 2 || !exists {
		t.Errorf("Failed to laod Data")
	}
}
