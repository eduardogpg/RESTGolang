package config

import(
    "html/template"
)

var Templates = template.Must(template.New("t").ParseGlob("../../templates/**/*.html") )
var ErrorTemplate = template.Must(template.ParseFiles("../../templates/error.html"))

func Initializer(){
  databaseInitializer()
}
