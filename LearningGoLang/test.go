package main
import (
  "fmt"
  "strconv"
)

//Types

type(
  Age int
  Name string
)

var age Age = 20
// Interesting things you can do
const Monday, Tuesday, Wednesday, Thursday, Friday int = 1,2,3,4,5

const Shit, Number, boool = "Poop", 4, false

type(
  Celsius float32
  Fahrenheit float32
)


// Basic Functions
func main() {
  

  var t Celsius
  var temp Fahrenheit
  temp = Fahrenheit(t + 2)

  // This is only applicable inside functions
  // Declaring new Variable
  variable := 24
  fmt.Println("Hello, World!")
  fmt.Println(age)
  fmt.Println(Shit, Number, boool)
  fmt.Println(variable)
  fmt.Println(temp)

  s := "Hel" + "lo"
  s += "Universe!"
  fmt.Println(s)

  if 5 > 10 {
    fmt.Println(s)
  } else if 10 > 5 {
    fmt.Println(age)
  } else {
    fmt.Println(Shit)
  }

  var num int
  var err error
  num, err = strconv.Atoi(s)
  
  if err != nil {
    fmt.Println("what the damn tis aint convertible", err)
  }

  fmt.Println(num)
  


}





