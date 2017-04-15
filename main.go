package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"./handlers"
	"./config"
)

func init(){
	
}

func main() {
	// defer models.CloseConnection()
	
	mux := mux.NewRouter()
	url := config.UrlServer()
	
	mux.HandleFunc("/api/v1/users/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	
	log.Println(url)
	log.Fatal(http.ListenAndServe(url, mux) )

}