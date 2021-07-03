package bank

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Tokyo",
			Phone:   "0120-123-456",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Tokyo, Japan",
			Phone:   "0120-123-456",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := account.Deposit(-10); err == nil {
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Error("statement doesn't have the proper format")
	}
}

func TestTransfer(t *testing.T) {
	accountA := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Tokyo, Japan",
			Phone:   "0120-123-456",
		},
		Number:  1001,
		Balance: 30,
	}
	accountB := Account{
		Customer: Customer{
			Name:    "Ryo",
			Address: "Saitama, Japan",
			Phone:   "0120-123-457",
		},
		Number:  1002,
		Balance: 0,
	}

	err := accountA.Transfer(50, &accountB)

	if accountA.Balance != 20 || accountB.Balance != 10 {
		t.Error("transfer from account A to account B is not working", err)
	}
}
