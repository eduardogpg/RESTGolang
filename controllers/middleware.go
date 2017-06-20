package controllers

import(
  "net/http"
  "../utils"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

//http://www.alexedwards.net/blog/making-and-using-middleware
func Authentication(next customeHandler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if !utils.IsAuthenticated(r){
      http.Redirect(w, r, "/users/login", http.StatusSeeOther)
      return
    }
    next(w, r)  
  })
}