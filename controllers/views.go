package controllers

import(
  "net/http"
  "html/template"
  "log"
)

var templates = template.Must(template.New("t").ParseGlob("/controllers/templates/**/*.html") )
var errorTemplate = template.Must(template.ParseFiles("/templates/error.html"))

func handlerError(w http.ResponseWriter, status int){
  w.WriteHeader(status)
  errorTemplate.Execute(w, nil)
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}){
  w.Header().Set("Content-Type", "text/html")
  err := templates.ExecuteTemplate(w, name, data)
  
  if err != nil{
    log.Println(err)
    handlerError(w, http.StatusInternalServerError)
  }
}
