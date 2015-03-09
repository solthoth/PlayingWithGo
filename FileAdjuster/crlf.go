package main

import(
  "fmt"
  "flag"
  "os"
  "strings"
  "bufio"
  "runtime"
)

var
  sCrLf string
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
  fmt.Println("Input file: "+sFile)
  if (strings.Trim(sFile," ")==""){
    sFile = "crlf"
  }
  fIn, errIn := os.Open(sFile)
  fOut, errOut := os.Create(sFile+".out")
  if (errIn == nil) || (errOut == nil){}
  defer fOut.Close();
  defer fIn.Close();
  fReader := bufio.NewReader(fIn)
  buffer := make([]byte,512)
  fmt.Print("Parsing file")
  var err error
  var bytes int
  for {
    bytes, err = fReader.Read(buffer)
    //fmt.Println(bytes)
    if bytes > 0 {
      fmt.Print(".")
      //fmt.Print(string(buffer))
      fOut.WriteString(string(buffer)+sCrLf)
      //break
    }
    if (err != nil){
      break//done working
    }
  }
  fmt.Println("Done!")
  fmt.Println("Output file: "+sFile+".out")
}

func init(){
  if runtime.GOOS == "windows" {
    sCrLf = "\r\n"
  } else {
    sCrLf = "\n"
  }
}
