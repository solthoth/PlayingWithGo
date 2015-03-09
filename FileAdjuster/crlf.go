package main

import(
  "fmt"
  "flag"
  "os"
  "strings"
)

func ParamStr(i int) string {
  if (i>=len(os.Args)) || (i<0) {
    return ""
  } else {
    return os.Args[i]
  }
}

func ReadBuffer(f *os.File, count int) string {
  buffer := make([]byte, count)
  f.Read(buffer)
  return string(buffer)
}

func main() {
  flag.Parse()
  sFile := ParamStr(1)
  fmt.Println(sFile)
  if (strings.Trim(sFile," ")==""){
    sFile = "crlf"
  }
  fIn, errIn := os.Open(sFile)
  fOut, errOut := os.Create(sFile+".out")
  if (errIn == nil) || (errOut == nil){}
  defer fOut.Close();
  defer fIn.Close();


}
