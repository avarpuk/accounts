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

	var (
		inputStr, change string
		selectedAccount  bool
		stop             = 0
	)

	for stop == 0 {
		fmt.Print(colorClear)

		accounts[Btoi(selectedAccount)].GetAccount()
		fmt.Print("\nAvailable commands:\n" +
			"1 - switches account\n" +
			"2 - changes the name of selected account\n" +
			"3 - changes the age of selected account\n" +
			"4 - stops the program\n" +
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
			fmt.Print("New name: ")
			if scanner.Scan() {
				change = scanner.Text()
			}
			accounts[Btoi(selectedAccount)].ChangeName(change)
		case "3":
			fmt.Print("\nNew name: ")
			if scanner.Scan() {
				change = scanner.Text()
			}
			newage, err := strconv.Atoi(change)
			if newage < 0 {
				fmt.Printf("%v is not available for pos ''Age''", newage)
				panic(err)
			}
			if err != nil {
				fmt.Printf("%v is not available for pos ''Age''", newage)
				panic(err)
			}
			accounts[Btoi(selectedAccount)].ChangeAge(newage)
		case "4":
			stop = 1
		default:
			fmt.Printf("%v is not an availabe comand\n", inputStr)
		}
	}
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
