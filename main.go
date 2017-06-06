package main

import(
  "log"
  "net/http"
  "./controllers"
  "github.com/gorilla/mux"
)

func main() {
  
  mux := mux.NewRouter()
  
  mux.Handle("/", controllers.Authentication(controllers.Index)).Methods("GET")
  mux.Handle("/login", controllers.UnAuthentication(controllers.Login)).Methods("GET", "POST")
  mux.Handle("/register", controllers.UnAuthentication(controllers.Register)).Methods("GET", "POST")
  mux.HandleFunc("/logout", controllers.Logout).Methods("GET", "POST")
  
  mux.HandleFunc("/api/v1/users/", controllers.GetUsers).Methods("GET")
  mux.HandleFunc("/api/v1/users/{id:[0-9]+}", controllers.GetUser).Methods("GET")
  mux.HandleFunc("/api/v1/users/", controllers.CreateUser).Methods("POST")
  mux.HandleFunc("/api/v1/users/{id:[0-9]+}", controllers.UpdateUser).Methods("PUT")
  mux.HandleFunc("/api/v1/users/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")
  
  log.Println("El servidor esta a la escucha en el puerto 3000")
  log.Fatal(http.ListenAndServe(":3000", mux) )

}