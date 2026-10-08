package main


// Part 4: Value vs pointer receivers

type Counter struct{ n int }

func (c Counter) IncValue() { c.n++	}
func (c *Counter) IncPointer() { c.n++ }

////////////////////////////////////////////////////////////////

type Account struct {
	Owner string
	Balance float64
}

func (a *Account) Deposit(amount float64) { a.Balance += amount }