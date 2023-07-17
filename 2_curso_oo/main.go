package main

import (
	"fmt"
    c "main/accounts"
    "main/customers"
    "main/utils"
)

func payBill(account verifyAccount, billValue float64) {
    account.CashWithdrawal(billValue)
}

type verifyAccount interface {
    CashWithdrawal(value float64) string
}

func main() {

    utils.GetDate()

    customerAlisson := customers.Holder{Name: "Alisson", CPF: "12312442142", Profession: "Analyst"}
    account := c.CheckingAccount{Holder: customerAlisson}
    account2 := c.CheckingAccount{Holder: customers.Holder{Name: "Jessica", CPF: "12389042142", Profession: "Analyst"}}

    account.DepositCash(900)
    println(account.CashWithdrawal(850))

    fmt.Println("saldo atual:", account.GetBalance())

    fmt.Println(account.DepositCash(100))

    fmt.Println(account.Transfer(100, &account2))

    payBill(&account, 10)

    fmt.Println("saldo atual conta alisson:", account.GetBalance())

    fmt.Println("saldo atual conta Jessica:", account2.GetBalance())

}