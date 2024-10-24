package main

import "fmt"


func sumInts(list ...int) (sum int){
  
  for _, v := range list {
    sum = sum + v
  }

  return


}
