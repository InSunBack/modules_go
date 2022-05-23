package account

import "fmt"

type account struct {
	owner   string
	balance int
}

func NewAccount(owner_ string) *account {
	account := account{owner: owner_, balance: 0}
	return &account
}

func PrintAccount(acc account) {
	fmt.Println("owner : ", acc.owner, " balance : ", acc.balance)
}

func Adder(a int, b int) {
	fmt.Println(a + b)
}
