package app

import(
  "github.com/gorilla/mux"
  "./handler"
)

func getRoutes() *mux.Router {
  mux := mux.NewRouter()
    
  mux.HandleFunc("/api/v1/users/", handler.GetUsers).Methods("GET")
  mux.HandleFunc("/api/v1/users/", handler.CreateUser).Methods("POST")
  
  mux.Handle("/api/v1/users/{id:[0-9]+}", handler.GetUser).Methods("GET")
  mux.Handle("/api/v1/users/{id:[0-9]+}", handler.UpdateUser).Methods("PUT")
  mux.Handle("/api/v1/users/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")
  
  // mux.Handle("/api/v1/users/{id:[0-9]+}", utils.MiddlewareFunc(handler.DeleteUser)).Methods("DELETE")
  return mux  
}