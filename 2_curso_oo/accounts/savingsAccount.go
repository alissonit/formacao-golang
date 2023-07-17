package accounts

import "main/customers"

type SavingsAccount struct {
	Holder customers.Holder
	AgencyNumber, AccountNumber, Operation int
	balance float64
}

func (c *SavingsAccount) DepositCash(depositCash float64) (string, float64) {
    if depositCash > 0 {
        c.balance += depositCash
        return "Deposito realizado com sucesso", c.balance
    } else {
        return "Valor do deposito menor que zero", c.balance
    }
}

func (c *SavingsAccount) CashWithdrawal(withDrawalValue float64) string {

    canWithDrawal := withDrawalValue > 0 && withDrawalValue <= c.balance

    if canWithDrawal {
        c.balance -= withDrawalValue
        return "saque realizado com sucesso"
    } else {
        return "saldo insuficiente"
    }

}

func (c *SavingsAccount) GetBalance() float64 {
    return c.balance
}
