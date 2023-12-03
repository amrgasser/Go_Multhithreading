package main

import (
	"Paytabs/proj/load"
	"Paytabs/proj/transfer"
	"Paytabs/proj/user"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// args (accouunt to pay the money: payerId, account to receiver: receiverId, amount as a string)
// method number 1 is to use locks
func main() {
	// Array containingg all users
	var users []user.User

	// HashMap containing all users with hash key ID to make lookup O(1) operation
	var userMap = make(map[string]int)

	//First on Run load the data to ingest into memory
	load.LoadData("data.json", &userMap, &users)
	fmt.Print("\nEnter your Command:\n")
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	ch := make(chan string, 10)

	for scanner.Scan() {
		// Grab Text from user input
		temp := scanner.Text()

		//Split the text with the first word as the command
		in := strings.Split(temp, " ")

		//Check commands and all
		switch in[0] {
		//list all users
		case "listall":
			fmt.Println(users)
			break
			//find a user using ID
		case "find":
			fmt.Println(userMap[in[1]])
			break
			//transfer credit format: transfer ${payerId} ${receiverId} ${amount}
		// case "listtransfers":
		// 	fmt.Println(transfers)
		case "transfer":
			if len(in) < 4 {
				fmt.Println("Could not transfer not enough variables")
				break
			}
			newTransfer := transfer.ParseInputAndReturnTransfer(in[1], in[2], in[3], &userMap, &users)
			// transfers = append(transfers, *transfer)
			for i := 0; i < 30; i++ {
				wg.Add(2)
				go transfer.PerformTransfer(newTransfer, &wg, ch)
			}

			wg.Wait()
			break
		default:
			fmt.Println("Invalid Command")
		}
		fmt.Println("Enter your Command:")
	}
}
