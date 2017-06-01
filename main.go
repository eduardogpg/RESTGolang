package main

import(
  "log"
  "net/http"
  "./controllers"
  "github.com/gorilla/mux"
)

func main() {
  
  mux := mux.NewRouter()
  
  mux.HandleFunc("/", controllers.Index).Methods("GET")

  mux.HandleFunc("/api/v1/users/", controllers.GetUsers).Methods("GET")
  mux.HandleFunc("/api/v1/users/{id:[0-9]+}", controllers.GetUser).Methods("GET")
  mux.HandleFunc("/api/v1/users/", controllers.CreateUser).Methods("POST")
  mux.HandleFunc("/api/v1/users/{id:[0-9]+}", controllers.UpdateUser).Methods("PUT")
  mux.HandleFunc("/api/v1/users/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")
  
  log.Println("El servidor esta a la escucha en el puerto 8000")
  log.Fatal(http.ListenAndServe(":8000", mux) )

}