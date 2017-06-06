package controllers

import(
  "log"
  _ "../utils"
  "net/http"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

//http://www.alexedwards.net/blog/making-and-using-middleware
func LoginMiddleware(next customeHandler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Before Handler!")
    next(w, r)
    log.Println("After Handler!")
  })
}