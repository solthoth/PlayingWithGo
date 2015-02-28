package main

import "fmt"
import "strconv"

func main(){
  var
    firstName, lastName string = "Carlos", "Barajas"
  age := 27

  fmt.Println("Hello, World!")
  fmt.Println("My name is " + firstName + " " + lastName)
  fmt.Println("I am " + IntToStr(age))
  fmt.Println("and I like tacos and burritos.")
  fmt.Println("")

  lastname := "Sanchez"
  fmt.Println("Hello, Ladies..")
  fmt.Println("My name is " + firstName + " " + lastname)
  fmt.Println("I am " + IntToStr(age))
  fmt.Println("and I too, like tacos and burritos ")
  fmt.Println("but I also enjoy piña coladas and long walks on the beach.")
}
func IntToStr(i int) string {
  return strconv.FormatInt(int64(i),10)
}
