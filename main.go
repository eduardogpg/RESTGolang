package main

import(
 "log"
  "net/http"
  "./controllers/"
  "./controllers/api/v1"
  "github.com/gorilla/mux"
)
  
func main(){

    mux := mux.NewRouter()
    mux.HandleFunc("/", controllers.Index)

    mux.HandleFunc("/users/new", controllers.NewUser).Methods("GET", "POST")
    mux.HandleFunc("/users/login", controllers.Login).Methods("GET", "POST")
    mux.HandleFunc("/users/logout", controllers.Logout).Methods("GET")
    mux.Handle("/users/edit", controllers.Authentication(controllers.EditUser)).Methods("GET")
    mux.Handle("/users/update", controllers.Authentication(controllers.UpdateUser)).Methods("PUT")
    
    mux.HandleFunc("/api/v1/users/", v1.GetUsers).Methods("GET")
    mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.GetUser).Methods("GET")
    mux.HandleFunc("/api/v1/users/", v1.CreateUser).Methods("POST")
    mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.UpdateUser).Methods("PUT")
    mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.DeleteUser).Methods("DELETE")
    
    assets := http.FileServer(http.Dir("assets"))
    statics := http.StripPrefix("/assets/", assets)
    mux.PathPrefix("/assets/").Handler(statics)

    log.Println("El servidor a la escucha en el puerto :3000")
    log.Fatal( http.ListenAndServe(":3000", mux))
}
