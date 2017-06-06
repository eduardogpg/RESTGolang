package controllers

import(
  "net/http"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

//http://www.alexedwards.net/blog/making-and-using-middleware
func Authentication(next customeHandler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if isAuthenticated(r){
      next(w, r)  
    }else{
      http.Redirect(w, r, "/login", http.StatusSeeOther)
    }
  })
}

func UnAuthentication(next customeHandler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if !isAuthenticated(r){
      next(w, r)  
    }else{
      http.Redirect(w, r, "/", http.StatusSeeOther)
    }
  })
}