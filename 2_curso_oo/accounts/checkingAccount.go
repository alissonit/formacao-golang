package accounts

import (
    "fmt"
    "main/customers"
)


type CheckingAccount struct {
    Holder       customers.Holder
    AgencyNumber int
    AccountNumber   int
    balance         float64
}

func createCheckingAccountV1() {
    accountUser := CheckingAccount{
        Holder: customers.Holder{Name: "Alisson", CPF: "1232132", Profession: "Analyst"},
        AgencyNumber: 123,
        AccountNumber: 1343433,
    }

    fmt.Println(accountUser)
}

func createCheckingAccountV2() {
    accountUser := CheckingAccount{customers.Holder{Name: "Alisson", CPF: "1232132", Profession: "Analyst"}, 123,1343433, 100.50,}

    fmt.Println(accountUser)
}

func CreateCheckingAccountV3(params CheckingAccount) CheckingAccount {
    var accountUser CheckingAccount

    accountUser.Holder = params.Holder
    accountUser.balance = params.balance

    fmt.Println(accountUser)

    return accountUser
}

func (c *CheckingAccount) DepositCash(depositCash float64) (string, float64) {
    if depositCash > 0 {
        c.balance += depositCash
        return "Deposito realizado com sucesso", c.balance
    } else {
        return "Valor do deposito menor que zero", c.balance
    }
}

func (c *CheckingAccount) CashWithdrawal(withDrawalValue float64) string {

    canWithDrawal := withDrawalValue > 0 && withDrawalValue <= c.balance

    if canWithDrawal {
        c.balance -= withDrawalValue
        return "saque realizado com sucesso"
    } else {
        return "saldo insuficiente"
    }

}

func (c *CheckingAccount) Transfer(transferValue float64, destAccount *CheckingAccount) bool {
    if transferValue < c.balance && transferValue > 0 {
        c.balance -= transferValue 
        destAccount.DepositCash(transferValue)
        return true
    } else {
        return false
    }
}

func (c *CheckingAccount) GetBalance() float64 {
    return c.balance
}
