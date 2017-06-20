package controllers

import(
  "fmt"
  "net/http"
  "../utils"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

//http://www.alexedwards.net/blog/making-and-using-middleware
func Authentication(next customeHandler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    
    if !utils.IsAuthenticated(r){
      fmt.Println("El usuario no esta atenticado!")
      http.Redirect(w, r, "/users/login", http.StatusSeeOther)
      return
    }

    next(w, r)  
  })
}

func UnAuthentication(next customeHandler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    next(w, r)  
  })
}

func CRSFToken(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.FormValue("CSRF-token") == ""{
      fmt.Println("CRSFToken exists")
      http.Redirect(w, r, r.URL.Path , http.StatusSeeOther)  
    }
    
    next.ServeHTTP(w, r)
  })
}
