package main

import(
  "log"
  "net/http"
  )

type String string

type Struct struct{
  Greetings string
  Punct     string
  Who       string
}

func main(){
  //doesnt work, am I missing something?
  http.Handle("/string",String("I'm a frayed knot."))
  http.Handle("/struct",&Struct{"Hello",":","Gophers!"})
  log.Fatal(http.ListenAndServe(":4000",nil))
}
