package main

import(
  "fmt"
  "io/ioutil"
  "net/http"
  "html/template"
  )

var templates = template.Must(template.ParseFiles("edit.html","view.html"))

type Page struct {
  Title string
  Body []byte
}

func (p *Page) Save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  fmt.Println("Calling loadPage...")
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
  /*t, err := template.ParseFiles(tmpl + ".html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err := t.Execute(w, p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }*/
  //replaces above
  err := templates.ExecuteTemplate(w, tmpl+".html", p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Calling viewHandler..." + r.URL.Path)
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  if (p != nil) {
    //fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
  } else {
    //fmt.Println("nil value returned: ", title)
    http.Redirect(w, r, "/edit/"+title, http.StatusFound)
    return
  }
}

func editHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Calling editHandler..." + r.URL.Path)
  title := r.URL.Path[len("/edit/"):]
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  /*fmt.Fprintf(w, "<h1>Editing %s</h1>"+
                 "<form action=\"/save/%s\" method=\"POST\""+
                 "<textarea name=\"body\">%s</textarea><br>"+
                 "<input type=\"submit\" value=\"Save\">"+
                 "</form>", p.Title, p.Title, p.Body)*/
  //replaces above
  //t, _ := template.ParseFiles("edit.html")
  //t.Execute(w, p)
  //replaces above
  renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request){
  title := r.URL.Path[len("/save/"):]
  body := r.FormValue("body")
  p := &Page{Title: title, Body: []byte(body)}
  err := p.Save()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
  /*p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
  p1.save()
  p2, _ := loadPage("TestPage")
  fmt.Println(string(p2.Body))*/
  fmt.Println("Starting server...")
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  http.HandleFunc("/save/", saveHandler)
  http.ListenAndServe(":8080", nil)
}
