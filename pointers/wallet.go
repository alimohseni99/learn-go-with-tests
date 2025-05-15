package pointers

type Wallet struct{}

func (w Wallet) Deposit(amount int) int{ 
	return amount
}

func (w Wallet) Balance() int{
	return 0
}