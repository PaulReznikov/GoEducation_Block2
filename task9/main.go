package main

import (
	"fmt"
	"time"
)

type BankAccount struct {
	Owner           string
	Balance         float64
	AccountType     string
	DataAccountOpen time.Time
	//AccruedInterest float64
	Transactions []Transaction
}

type Transaction struct {
	Operation               string
	amount                  float64
	BalanceAfterTransaction float64
	DataOperation           time.Time
	//To                      BankAccount
}

func (ba *BankAccount) Deposit(amount float64) {
	time.Sleep(2 * time.Second)
	ba.Balance += amount
	ba.Transactions = append(ba.Transactions, Transaction{
		Operation:               "Deposit",
		amount:                  amount,
		BalanceAfterTransaction: ba.Balance,
		DataOperation:           time.Now(),
		//To:                      *ba,
	})
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if ba.AccountType == "Checking" {
		if ba.Balance-(amount+amount*0.01) > 0 {
			time.Sleep(2 * time.Second)
			ba.Balance -= amount + amount*0.01
			ba.Transactions = append(ba.Transactions, Transaction{
				Operation:               "Withdraw",
				amount:                  amount + amount*0.01,
				BalanceAfterTransaction: ba.Balance,
				DataOperation:           time.Now(),
				//To:                      BankAccount{},
			})
			return nil
		}
	} else if ba.Balance-amount > 0 {
		time.Sleep(2 * time.Second)
		ba.Balance -= amount
		ba.Transactions = append(ba.Transactions, Transaction{
			Operation:               "Withdraw",
			amount:                  amount,
			BalanceAfterTransaction: ba.Balance,
			DataOperation:           time.Now(),
			//To:                      BankAccount{},
		})
		return nil
	}

	return fmt.Errorf("списание с аккаунта пользователя %v невозможно,"+
		" на балансе недостаточно средств, текущий баланс: %v", ba.Owner, ba.Balance)
}

func (ba *BankAccount) GetBalance() {
	fmt.Printf("Баланс: %v\n", ba.Balance)
}

func (ba *BankAccount) Transfer(to *BankAccount, amount float64) error {
	if ba.AccountType == "Locked" {
		return fmt.Errorf("аккаунт пользователя %v - заблокирован, транзакция невозможна", ba.Owner)
	}

	err := ba.Withdraw(amount)
	if err != nil {
		return err
	}

	to.Deposit(amount)
	//ba.Transactions[len(ba.Transactions)-1].To = *to
	return nil
}

func (ba *BankAccount) CalculateInterest() float64 {
	if ba.AccountType == "Saving" {
		percent := ba.Balance*1.18/12 - ba.Balance
		ba.Balance *= 1.18 / 12
		return percent
	} else if ba.AccountType == "Checking" {
		percent := ba.Balance*1.11/12 - ba.Balance
		ba.Balance *= 1.11 / 12
		return percent
	}

	return 0
}

// GenerateStatement /////////////////////////////////////////////////////
func (ba *BankAccount) GenerateStatement(begin, end time.Time) ([]Transaction, error) {
	periodTransaction := make([]Transaction, 0)

	if ba.DataAccountOpen.Before(end) && begin.Before(end) && !end.After(time.Now()) { //??? time.Date()
		for _, transaction := range ba.Transactions {
			if !transaction.DataOperation.Before(begin) && !transaction.DataOperation.After(end) {
				periodTransaction = append(periodTransaction, transaction)
			}
		}

		return periodTransaction, nil
	}

	return []Transaction{}, fmt.Errorf("указан некорректный временой промежуток")
}

func main() {

	bankAcc := BankAccount{
		Owner:           "Reznikov",
		Balance:         0,
		AccountType:     "Checking",
		DataAccountOpen: time.Date(2017, time.October, 14, 0, 0, 0, 0, time.UTC),
		Transactions:    nil,
	}

	beginPeriod := time.Now()

	bankAcc.GetBalance()
	bankAcc.Deposit(1000)
	bankAcc.GetBalance()
	bankAcc.Deposit(500)
	bankAcc.GetBalance()
	bankAcc.Withdraw(500)
	bankAcc.GetBalance()

	endPeriod := time.Now()

	//beginPeriod := time.Date(2024, time.October, 24, 14, 59, 1, 1, time.UTC)

	extract, err := bankAcc.GenerateStatement(beginPeriod, endPeriod)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range extract {
			fmt.Printf("%v\n", val)
		}
	}
}

//bankAcc := BankAccount{
//	Owner:           "Reznikov Pavel",
//	Balance:         1000,
//	AccountType:     "Checking",
//	DataAccountOpen: time.Date(2017, time.October, 14, 0, 0, 0, 0, time.UTC),
//	Transactions: []Transaction{
//		{
//			Operation:     "Deposit",
//			amount:        5000,
//			DataOperation: time.Date(2017, time.October, 20, 0, 0, 0, 0, time.UTC),
//			To:            BankAccount{},
//		},
//		{
//			Operation:     "Deposit",
//			amount:        5000,
//			DataOperation: time.Date(2020, time.October, 15, 0, 0, 0, 0, time.UTC),
//			To:            BankAccount{},
//		},
//		{
//			Operation:     "Deposit",
//			amount:        5000,
//			DataOperation: time.Date(2021, time.August, 2, 0, 0, 0, 0, time.UTC),
//			To:            BankAccount{},
//		},
//		{
//			Operation:     "Deposit",
//			amount:        5000,
//			DataOperation: time.Date(2024, time.October, 5, 0, 0, 0, 0, time.UTC),
//			To:            BankAccount{},
//		},
//	},
//}
