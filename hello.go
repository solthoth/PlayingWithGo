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
}

func IntToStr(i int) string {
  return strconv.FormatInt(int64(i),10)
}
