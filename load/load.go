package load

import (
	"Paytabs/proj/user"
	"encoding/json"
	"io/ioutil"
	"log"
)

// Load the data from the above method to users array
func LoadData(filepath string, userMap *map[string]int, users *[]user.User) {

	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	err = json.Unmarshal(content, &users)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for index, user := range *users {
		(*userMap)[user.Id] = index
	}
}
