package controllers

import(
  "net/http"
  "../utils"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

//http://www.alexedwards.net/blog/making-and-using-middleware
func Authentication(next customeHandler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if utils.IsAuthenticated(r){
      next(w, r)  
    }else{
      http.Redirect(w, r, "/login", http.StatusSeeOther)
    }
  })
}

func UnAuthentication(next customeHandler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if !utils.IsAuthenticated(r){
      next(w, r)  
    }else{
      http.Redirect(w, r, "/", http.StatusSeeOther)
    }
  })
}

func CRSFToken(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.FormValue("CSRF-token") == ""{
      http.Redirect(w, r, r.URL.Path , http.StatusSeeOther)  
    }
    
    next.ServeHTTP(w, r)
  })
}
