package app

import(
  "net/http"
)

func CreateServer() *http.Server{
  mux := getRoutes()
  
  server := &http.Server{
    Addr: "localhost:3000",
    Handler:  mux,
  }
  
  return server
}