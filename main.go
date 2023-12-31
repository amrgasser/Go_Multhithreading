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
func main() {
	// Array containingg all users
	var users []user.User
	// HashMap containing all users with hash key ID to make lookup O(1) operation
	var userMap = make(map[string]int)
	var transfers []transfer.Transfer

	//First on Run load the data to ingest into memory
	load.LoadData("data.json", &userMap, &users)

	scanner := bufio.NewScanner(os.Stdin)

	var wg sync.WaitGroup
	ch := make(chan string, 10)

	fmt.Print("\nEnter your Command:\n")
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
		case "listtransfers":
			fmt.Println(transfers)
			//transfer credit format: transfer ${payerId} ${receiverId} ${amount}
		case "transfer":
			if len(in) < 4 {
				fmt.Println("Could not transfer not enough variables")
				break
			}
			newTransfer := transfer.ParseInputAndReturnTransfer(in[1], in[2], in[3], &userMap, &users)
			wg.Add(2)
			go transfer.PerformTransfer(newTransfer, &wg, ch)
			break
		default:
			fmt.Println("Invalid Command")
		}
		fmt.Println("Enter your Command:")
	}
}
