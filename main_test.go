package main

import (
	"Paytabs/proj/load"
	"Paytabs/proj/transfer"
	"Paytabs/proj/user"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestMain(t *testing.T) {
	var users []user.User
	var userMap = make(map[string]int)
	var transfers []transfer.Transfer

	var wg sync.WaitGroup
	ch := make(chan string, 10)

	load.LoadData("./data.json", &userMap, &users)

	file, err := os.Open("./commands.txt")
	if err != nil {
		t.Errorf("Error opening file: %s", err)
		return
	}
	rd := bufio.NewReader(file)

	// Read and print each line
	for {
		line, err := rd.ReadString('\n')
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				break
			}

			t.Errorf("read file line error: %v", err)
			return
		}

		in := strings.Split(strings.Split(line, "\n")[0], " ")
		fmt.Sprintln(in[3])
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
				t.Errorf("Could not transfer not enough variables")
				break
			}

			newTransfer := transfer.ParseInputAndReturnTransfer(in[1], in[2], in[3], &userMap, &users)
			wg.Add(2)
			go transfer.PerformTransfer(newTransfer, &wg, ch)
			wg.Wait()
			break
		default:
			t.Errorf("Invalid Command")
		}
	}
	// wg.Wait()
	file.Close()
}
