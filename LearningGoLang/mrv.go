package main

import (
  "fmt"

)

func SumProductDiff(i, j int) (s int, p int, d int){

  return i + j, i * j, i - j

}

func main() {

  res, _ := SumProductDiff(3,4)

  fmt.Println(res)

}
