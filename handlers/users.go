package handlers

import(
	"errors"
	"strconv"
	"net/http"
	"encoding/json"
	
	"../models"
	"github.com/gorilla/mux"
)

// curl -X GET http://localhost:8000/api/v1/users/
func GetUsers(w http.ResponseWriter, r *http.Request){
	models.SendData(w, models.GetUsers() )
}

// curl -X GET http://localhost:8000/api/v1/users/1
func GetUser(w http.ResponseWriter, r *http.Request){
	if user, err := getUserByRequest(r); err != nil{
		SendNotFound(w)
	}else{
		SendData(w, user)
	}
}

// curl -X POST http://localhost:8000/api/v1/users/ -d '{"username" : "eduardo", "password": "password", "email": "eduardo@codigofacilito.com"}' -H "Content-type: application/json"
func CreateUser(w http.ResponseWriter, r *http.Request){
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil{
		models.SendUnprocessableEntity(w)
	}else{
		user.SetPassword(user.Password)
		
		if err := user.Save(); err != nil{
			models.SendBadRequest(w)
		}else{
			SendData(w, user)	
		}
	}
}

// curl -X PUT http://localhost:8000/api/v1/users/1 -d '{"username" : "new_eduardo", "new_password": "password", "email": "new_eduardo@codigofacilito.com"}' -H "Content-type: application/json"
func UpdateUser(w http.ResponseWriter, r *http.Request){
	user, err := getUserByRequest(r)
	
	if err != nil{
		SendNotFound(w)
		return
	}

	userRequest := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userRequest); err != nil{
		SendUnprocessableEntity(w)
		return
	}

	user.Username = userRequest.Username
	user.SetPassword(userRequest.Password)
	user.Email = userRequest.Email
	
	//user.Valid()?
	if err := user.Save(); err != nil{
		models.SendBadRequest(w)
	}else{
		SendData(w, user)	
	}
}

// curl -X DELETE http://localhost:8000/api/v1/users/1 
func DeleteUser(w http.ResponseWriter, r *http.Request){
	if user, err := getUserByRequest(r); err != nil{
		SendNotFound(w)
	}else{
		
		if err := user.Delete(); err != nil{
			models.SendBadRequest(w)
		}else{
			SendNoContent(w)	
		}
	}
}

func getUserByRequest(r *http.Request) (*models.User, error){
	vars := mux.Vars(r)
	id, _ :=  strconv.Atoi( vars["id"] )
	
	if user := models.GetUser("id", id); user.Id <= 0{
		return user, errors.New("User not found")
	}else{
		return user, nil
	}
}
