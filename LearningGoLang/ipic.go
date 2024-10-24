package main

import "fmt"

// mutual recursion pyutanginang to

func isPrime(num int) bool{
  if num <= 1{ 
    return false
  }
  return !isComposite(num)

}
func isComposite(num int) bool{
  if num <= 1{ 
    return false
  }   
  for i := 2; i*i <= num; i++{
    if num % i == 0{
      return true
    } 
  }
  return false
}

func main(){

}
