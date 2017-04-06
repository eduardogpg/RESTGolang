package main

import(
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "./handler"
  "./models"
  "./utils"
  "./config"
)

func main() {
  models.SetDefaultUser()
  config.Initializer()
  
  mux := mux.NewRouter()

  mux.HandleFunc("/login/", handler.GetUsers).Methods("GET", "POST")
  
  
  mux.HandleFunc("/api/v1/users/", handler.GetUsers).Methods("GET")
  mux.HandleFunc("/api/v1/users/", handler.CreateUser).Methods("POST")
  
  mux.Handle("/api/v1/users/{id:[0-9]+}", utils.MiddlewareFunc(handler.GetUser)).Methods("GET")
  mux.Handle("/api/v1/users/{id:[0-9]+}", utils.MiddlewareFunc(handler.UpdateUser)).Methods("PUT")
  mux.Handle("/api/v1/users/{id:[0-9]+}", utils.MiddlewareFunc(handler.DeleteUser)).Methods("DELETE")
  
  log.Println("Server listening port :3000")
  log.Fatal(http.ListenAndServe(":3000", mux))
}
