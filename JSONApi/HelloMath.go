package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  )

type MathType struct {
  Value1 int
  Value2 int
  Result int
}

func addHandler(httpWriter http.ResponseWriter, request *http.Request) {
  fmt.Println(request.URL.Path)
  input1 := 5
  input2 := 10
  result := input1 + input2
  jsonMath:=MathType{ Value1: input1, Value2: input2, Result: result }
  json.NewEncoder(httpWriter).Encode(jsonMath)
}

func main() {
  fmt.Println("Starting service...")
  http.HandleFunc("/MathResult/Add",addHandler)
  http.ListenAndServe(":8080", nil)
}
