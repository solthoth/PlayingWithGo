package main

import (
  "fmt"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  fmt.Printf("Program start\n")
  f, err := os.Create("temp.txt")
  check(err)
  //defer f.WriteString("This will not work\n")//While this does execute, it will not write to the file
  defer f.Close()//this causes the Close function to execute
                 //at the end of the code block
  defer f.WriteString("This will be written before the file closes\n")
  f.WriteString("This is only a test\n")
  defer f.WriteString("Where is this now?\n")
  f.WriteString("Here is more text\n")


  fmt.Printf("Program end\n")
}
