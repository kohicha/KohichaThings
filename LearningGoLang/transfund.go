package main

import "fmt"

// Defining the structure for an account
type Account struct {
    ID      int
    Balance float64
}

// Implement the function to transfer funds between accounts
func transferFunds(amount float64, source *Account, destination *Account) {
   // Write your implementation here
  if source.Balance < amount {
    goto badtransact
  } else{
    source.Balance = source.Balance - amount
    destination.Balance = amount + destination.Balance
    goto goodtransact
  }
  
  badtransact:
  fmt.Println("Transaction Failed, Insufficient Balance.")
  return

  goodtransact:
  fmt.Println("Transaction Successful!")
  
}

func main() {
    // Creating two accounts with initial balances
    account1 := Account{ID: 1, Balance: 100.0}
    account2 := Account{ID: 2, Balance: 50.0}

    // Demonstrating various transactions
    transferFunds(50.0, &account1, &account2)
    transferFunds(200.0, &account1, &account2)
}
