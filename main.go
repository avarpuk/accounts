package main

import (
	"bufio"
	"fmt"
	"learning/account"
	"os"
	"strconv"
	"strings"
)

const colorClear = "\033[H\033[2J"

func main() {
	
	accounts := []account.Account{
		&account.FirstAccount{},
		&account.SecondAccount{},
	}

	for {
		fmt.Print(colorClear)

		var (
			inputStr, change string
			selectedAccount  bool
		)

		accounts[Btoi(selectedAccount)].GetAccount()
		fmt.Print("\nAvailable commands:\n" +
			"1 - switches account\n" +
			"2 - gets all data of account\n" +
			"3 - changes the name of selected account\n" +
			"4 - changes the age of selected account\n" +
			"5 - stops the program\n" +
			"Selected command: ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			inputStr = strings.ToLower(scanner.Text())
		}

		switch inputStr {
		case "1":
			selectedAccount = !selectedAccount
			fmt.Printf("Account has successfully changed to %v", Btoi(selectedAccount))
		case "2":
			accounts[Btoi(selectedAccount)].GetAccount()
		case "3":
			change = scanner.Text()
			accounts[Btoi(selectedAccount)].ChangeName(change)
		case "4":
			change = scanner.Text()
			newage, err := strconv.Atoi(change)
			if newage <= 0 {
				fmt.Printf("%v is not available for pos ''Age''", newage)
				panic(err)
			}
			if err != nil {
				fmt.Printf("%v is not available for pos ''Age''", newage)
				panic(err)
			}
			accounts[Btoi(selectedAccount)].ChangeAge(newage)
		case "5":
			break
		default:
			fmt.Printf("%v is not an availabe comand", inputStr)
		}
	}
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
