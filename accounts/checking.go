package accounts

import "bank/clients"

type Checking struct {
	Holder       clients.Holder
	AgencyNumber int
	Number       int
	balance      float64
}

func (c *Checking) Withdraw(value float64) string {
	canWithdraw := value <= c.balance && value > 0

	if canWithdraw {
		c.balance -= value

		return "Saque realizado"
	}

	return "Saldo insuficiente"
}

func (c *Checking) Deposit(value float64) (string, float64) {
	isPositive := value > 0

	if isPositive {
		c.balance += value

		return "Saldo depositado", c.balance
	}

	return "Valor de deposito menor que zero", c.balance
}

func (c *Checking) Transfer(value float64, toAccount *Checking) bool {
	hasPositiveBalance := c.balance >= value && value > 0

	if hasPositiveBalance {
		c.balance -= value
		toAccount.balance += value

		return true
	}

	return false
}

func (c *Checking) GetBalance() float64 {
	return c.balance
}
