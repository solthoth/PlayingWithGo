package main

import(
  "fmt"
  "io/ioutil"
  "os"
  //"bufio"
  )

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func ReadFile(path string) string {
  buffer, err := ioutil.ReadFile(path)
  check(err)
  return string(buffer)
}

func ReadBuffer(f *os.File, count int) string {
  buffer := make([]byte, count)
  //Version A
  //bytesRead, err := f.Read(buffer)
  //check(err)
  //if bytesRead=0 {}
  //Version B
  //_, err := f.Read(buffer)
  //check(err)
  //Version C
  f.Read(buffer)
  return string(buffer)
}

func ReadLine(f *os.File) string {
  /* This version requires a global reader instead
  of using a global file handle
  reader := bufio.NewReader(f)
  buffer, _ := reader.ReadBytes('\n')
  return string(buffer)*/
  buffer := ""
  buff   := ""
  for ; buff != "\n"; {
    buff = ReadBuffer(f,1)
    buffer = buffer + buff
  }
  return buffer
}

func main() {
  var
    sFile string
  sFile = "README.md"
  fmt.Println("--- Reading all at once ---")
  file := ReadFile(sFile)
  fmt.Print(file)
  fmt.Println("---------------------------")

  fmt.Println("--- Reading a few bytes ---")
  fIn, err := os.Open(sFile)
  check(err)
  fmt.Println(ReadBuffer(fIn,10))
  fmt.Println(ReadBuffer(fIn,10))
  fIn.Close()
  fmt.Println("---------------------------")

  fmt.Println("------ Reading lines ------")
  fIn, err = os.Open(sFile)
  check(err)
  fmt.Print(ReadLine(fIn))
  fmt.Print(ReadLine(fIn))
  fmt.Print(ReadLine(fIn))
  fIn.Close()
  fmt.Println("---------------------------")
}
