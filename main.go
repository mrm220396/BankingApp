package main

import (
	"fmt"
	"time"
)

func DayHour() string {
	now := time.Now().Hour()

	if now >= 18 {
		return "Evening"
	} else if now >= 12 && now <= 17 {
		return "Afternoon"
	}

	return "Morning"
}

func menu() {
	var option int
	var quit bool = false
	var account AccountBank

	account.AccountAccess()

	for !quit {

		fmt.Printf(`Good %s %s %s
1] Deposit
2] Withdraw
3] Check Balance
4] Exit
Choose an option: `, DayHour(), account.FirstName, account.LastName)

		fmt.Scan(&option)

		switch option {
		case 1:
			account.Deposit()
		case 2:
			_, err := account.Withdraw()
			if err != nil {
				fmt.Println(err)
			}
			break
		case 3:
			account.CheckBalance()
			break
		case 4:
			quit = true
			fmt.Println("Thanks for choosing BankAPP =] ")
			break
		default:
			fmt.Println("Write a valid option")
			break
		}
	}

}

func main() {
	menu()
}
