package main

import (
	"errors"
	"fmt"
)

// Account represents a basic bank account structure
type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

// SavingsAccount embeds Account and adds an interest rate
type SavingsAccount struct {
	Account
	InterestRate float64
}

// CheckingAccount embeds Account and adds an overdraft limit
type CheckingAccount struct {
	Account
	OverdraftLimit float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid deposit amount")
	}

	a.Balance += amount
	fmt.Printf("Deposit of %.2f successful. New balance: %.2f\n", amount, a.Balance)
	return nil
}

// Withdraw subtracts an amount from the account balance if funds are sufficient
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid withdrawal amount")
	}
	if amount > a.Balance {
		return fmt.Errorf("Insufficent funds in %s. Balance: $%2f, Tried to withdraw: $%.2f", a.AccountNumber, a.Balance, amount)
	}
	a.Balance -= amount
	fmt.Printf("Withdrawal of $%.2f successful. New balance: $%.2f\n", amount, a.Balance)
	return nil
}

// GetBalance returns the current balance of the account
func (a *Account) GetBalance() float64 {
	return a.Balance
}

// String returns a formatted string representation of the account
func (a *Account) String() string {
	return fmt.Sprintf("Account Number: %s, Owner: %s, Balance: $%.2f", a.AccountNumber, a.OwnerName, a.Balance)
}

// ApplyInterest calculates and adds interest to the savings account balance
func (sa *SavingsAccount) ApplyInterest() {
	interest := sa.Balance * sa.InterestRate
	sa.Balance += interest
	fmt.Printf("Interest applied to %s. New balance: $%.2f\n", sa.AccountNumber, sa.Balance)
	err := sa.Deposit(interest)
	if err != nil {
		fmt.Println("ApplyInterest: Error depositing interest:", err)
	}

}

// Withdraw overrides the base Account Withdraw to allow for overdraft limits
func (ca *CheckingAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid withdrawal amount")
	}

	//Allow withdrawal up to Balance + overdraft limit
	if amount > ca.Balance+ca.OverdraftLimit {
		return fmt.Errorf("Insufficient funds in %s. Balance: $%.2f, Overdraft Limit: $%.2f, Tried to withdraw: $%.2f", ca.AccountNumber, ca.Balance, ca.OverdraftLimit, amount)
	}

	ca.Balance -= amount
	fmt.Printf("Withdrawal from overdraft account of $%.2f successful. New balance: $%.2f\n", amount, ca.Balance)
	return nil


}

func main() {

	fmt.Println("==========BANK ACCOUNT SYSTEM===========");

	// Initialize a Savings Account
	savingAcct := SavingsAccount{
		Account: Account{
			AccountNumber: "SAV-001",
			Balance:       1000.00,
			OwnerName:     "Alice",
		},
		InterestRate: 0.05,
	}

	fmt.Println("\n============== SAVINGS ACCOUNT OPERATION==========")

	fmt.Println(savingAcct.String())

	err := savingAcct.Deposit(500)
	if err != nil {
		fmt.Println("Error depositing to savings account:", err)
	}

	savingAcct.ApplyInterest()

	err = savingAcct.Withdraw(200)
	if err != nil {
		fmt.Println("Error withdrawing from savings account:", err)
	}


	fmt.Println("\nFINAL SAVINGS ACCOUNT DETAILS:", savingAcct.String())


		fmt.Println("\n============== OVERDRAFT ACCOUNT OPERATION==========")


	// Initialize a Checking Account with an overdraft limit
	checkingAcct := CheckingAccount{
		Account: Account{
			AccountNumber: "CHK-001",
			Balance:       2000.00,
			OwnerName:     "Bob",
		},
		OverdraftLimit: 170.00,
	}

	fmt.Println(checkingAcct.String())

	// Attempt a withdrawal that exceeds balance but stays within overdraft limit
	err = checkingAcct.Withdraw(2200)
	if err != nil {
		fmt.Println("Error withdrawing from checking account:", err)
	}

	fmt.Println("\nFINAL CHECKING ACCOUNT DETAILS:", checkingAcct.String())




}


//you can make it more better using interface and struct embedding to avoid code repetition and make it more extensible. For example, you could define a common interface for accounts and implement it for both savings and checking accounts. This way, you can easily add new account types in the future without modifying existing code.
