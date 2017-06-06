package controllers

import(
  "fmt"
  "strings"
  "net/http"
  "../utils"
  "../models"
)

func Index(w http.ResponseWriter, r *http.Request){
  login := IsLoggendIn(r)
  fmt.Println(login)
  
  utils.RenderTemplate(w, "index", nil)
}

func Restricted(w http.ResponseWriter, r *http.Request){
  utils.RenderTemplate(w, "restricted", nil)
}

func Register(w http.ResponseWriter, r *http.Request){
  context := make(map[string]interface{})
  context["User"] = models.User{}

  if r.Method == "POST"{
    //https://golang.org/pkg/net/http/#Request.ParseForm
    username := r.FormValue("username") //r.Form["username"][0]
    password := r.FormValue("password")
    email := r.FormValue("email")

    user, err := models.CreateUser(username, password, email)
    if err != nil{
      context["User"] = user
      context["Error"] = err.Error()
    }else{
      
    }
  }
  utils.RenderTemplate(w, "register", context)
}

func Login(w http.ResponseWriter, r *http.Request){
  if r.Method == "POST"{
    r.ParseForm()
    username := r.Form["username"][0]
    password := r.Form["password"][0]
    
    if err := models.Login(username, password); err != nil{
      message := utils.Message{err.Error(), "danger"}
      utils.RenderTemplate(w, "login", message)
    }else{
      fmt.Println("Generando cookie...")
      cookie := LoginCookie(username)
      http.SetCookie(w, &cookie)
      http.Redirect(w, r, "/restricted", http.StatusSeeOther)
    }
    return
  }
  
  utils.RenderTemplate(w, "login", nil)
}

func Logout(w http.ResponseWriter, r *http.Request){
  http.Redirect(w, r, "/login/", http.StatusSeeOther)
}


func LoginCookie(username string) http.Cookie {
  cookieValue := fmt.Sprintf("%s:%d", username, 100)
  return http.Cookie{Name: "SessionID", Value: cookieValue, HttpOnly: true}
}

func IsLoggendIn(r *http.Request) bool{
  cookie, err := r.Cookie("SessionID")
  if err != nil{
    return false
  }

  SessionID := cookie.Value
  fmt.Println(SessionID)
  array := strings.Split(SessionID, ":")
  email := array[0]
  SessionID = array[1]

  fmt.Println(email)
  fmt.Println(SessionID)

  return SessionID != ""
}







