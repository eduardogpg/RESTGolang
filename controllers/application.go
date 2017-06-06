package controllers

import(
  "net/http"
  "../utils"
  "../models"
)

func Index(w http.ResponseWriter, r *http.Request){
  utils.RenderTemplate(w, "index", nil)
}

func Restricted(w http.ResponseWriter, r *http.Request){
  utils.RenderTemplate(w, "restricted", nil)
}

func Register(w http.ResponseWriter, r *http.Request){
  context := make(map[string]interface{})
  context["User"] = models.User{}

  if r.Method == "POST"{ //https://golang.org/pkg/net/http/#Request.ParseForm
    username := r.FormValue("username") //r.Form["username"][0]
    password := r.FormValue("password")
    email := r.FormValue("email")

    if user, err := models.CreateUser(username, password, email); err != nil{
      context["User"] = user
      context["Error"] = err.Error()
    }else{
      http.Redirect(w, r, "/restricted", http.StatusSeeOther)
      return
    }
  }
  utils.RenderTemplate(w, "register", context)
}

func Login(w http.ResponseWriter, r *http.Request){
  context := make(map[string]interface{})
  
  if r.Method == "POST"{
    err := models.Login(r.FormValue("username"), r.FormValue("password"))
    if err != nil{
      context["Error"] = err.Error()
    }else{
      setSession(w, r)
      http.Redirect(w, r, "/", http.StatusSeeOther)
      return
    }
  }
  utils.RenderTemplate(w, "login", context)
}

func Logout(w http.ResponseWriter, r *http.Request){
  deleteSession(w)
  http.Redirect(w, r, "/login", http.StatusSeeOther)
}
