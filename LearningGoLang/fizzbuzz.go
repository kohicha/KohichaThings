package main

import (
  "fmt"
)

func main(){
  for i := 1; i <= 60; i++{
    if i % 3 == 0 && i % 5 == 0{
      fmt.Println("FizzBuzz")
      continue
    } 
    if i % 3 == 0 {
      fmt.Println("Fizz")
      continue
    }
    if i % 5 == 0 {
      fmt.Println("Buzz")
      continue
    }
    fmt.Println(i)
  }
}

/*
Fizz buzz
loop through numbers
if number is divisible by 3
print fizz
if number is divisible by 5
print buzz
if number is divisible by both
print fizzbuzz
else
print the number
*/
