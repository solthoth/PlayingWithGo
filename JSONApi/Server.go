package main

import(
  "fmt"
  //"html"
  "log"
  "net/http"

  "github.com/gorilla/mux"
  )

/*

func handler(w, htp.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080",nil)
}
*/

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)
  router.HandleFunc("/todos", TodoIndex)
  router.HandleFunc("/todos/{todoId}", TodoShow)
  log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  todoId := vars["todoId"]
  fmt.Fprintln(w, "Todo show:", todoId)
}
