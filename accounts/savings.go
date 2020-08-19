package accounts

import (
	"bank/clients"
)

type Savings struct {
	Holder       clients.Holder
	AgencyNumber int
	Number       int
	Operation    int
	balance      float64
}

func (c *Savings) Withdraw(value float64) string {
	canWithdraw := value <= c.balance && value > 0

	if canWithdraw {
		c.balance -= value

		return "Saque realizado"
	}

	return "Saldo insuficiente"
}

func (c *Savings) Deposit(value float64) (string, float64) {
	isPositive := value > 0

	if isPositive {
		c.balance += value

		return "Saldo depositado", c.balance
	}

	return "Valor de deposito menor que zero", c.balance
}

func (c *Savings) GetBalance() float64 {
	return c.balance
}
