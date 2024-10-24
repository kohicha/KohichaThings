package main

import "fmt"

// Lessen the code used
// Use Stricter type declaration
// Make it more readable


// Global variable
// var total int

// Create Constant Reusable variables
const (
  Factor1 = 2
  Factor2 = 3  
)

func main() {
  // local variable
    // Block level local variable
    //instead of explicitly declaring 2, we should create a reusable variable
    
  sum1 := calculateSum(5, Factor1)
  fmt.Println("Sum:", sum1)
    // Instead of repeatedly declaring mutable values, just declare it man.
    // count := 10
  sum2 := calculateSum(10, Factor2)
  fmt.Println("Sum:", sum2)
  total := sum1 + sum2
  fmt.Println("Total:", total)
}

func calculateSum(count, factor int) int {
    // Accessing local variable
    return count * factor
}
