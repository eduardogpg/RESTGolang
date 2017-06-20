package controllers

import(
  "net/http"
  "../utils"
  "../models"
)

func NewUser(w http.ResponseWriter, r *http.Request){
  context := make(map[string]interface{})
  context["User"] = models.User{}

  if r.Method == "POST"{
    //https://golang.org/pkg/net/http/#Request.ParseForm
    username := r.FormValue("username") //r.Form["username"][0]
    password := r.FormValue("password") //https://golang.org/pkg/net/http/#Request.FormValue
    email := r.FormValue("email")
    
    if user, err := models.CreateUser(username, password, email); err != nil{
      context["User"] = user
      context["Error"] = err.Error()
    }else{
      utils.SetSession(user, w, r)
      http.Redirect(w, r, "/", http.StatusSeeOther)
      return
    }
  }
  utils.RenderTemplate(w, "users/new", context)
}

func EditUser(w http.ResponseWriter, r *http.Request){
  context := make(map[string]interface{})
  user := utils.GetUser(r)
  context["User"] = user
  utils.RenderTemplate(w, "users/edit", context)
}

func Login(w http.ResponseWriter, r *http.Request){
  context := make(map[string]interface{})

  if r.Method == "POST"{
    user, err := models.Login(r.FormValue("username"), r.FormValue("password"))
    if err != nil{
      context["Error"] = err.Error()
    }else{
      createCookie(w)
      utils.SetSession(user, w, r)
      http.Redirect(w, r, "/", http.StatusSeeOther)
      return
    }
  }
  utils.RenderTemplate(w, "users/login", context)
}

func Logout(w http.ResponseWriter, r *http.Request){
  utils.DeleteSession(w)
  deleteCookie(w)
  http.Redirect(w, r, "/users/login", http.StatusSeeOther)
}

func createCookie(w http.ResponseWriter){
  cookie := &http.Cookie{
    Name: "cookie_name",
    Value: "slipknot!",
    Path: "/",
  }//You can only store about 4kb of data in a cookie

  //Expires: expire
  //time.Now().AddDate( year int, month int, day int)
  http.SetCookie(w, cookie)
}

func deleteCookie(w http.ResponseWriter){
  cookie := &http.Cookie{
    Name: "cookie_name",
    Value: "",
    MaxAge: -1,
  }
  http.SetCookie(w, cookie)
}



