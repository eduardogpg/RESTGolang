package controllers

import(
	"strconv"
	"errors"
	"net/http"
	"../utils"
	"../models"
	"encoding/json"
	"github.com/gorilla/mux"
)

// curl -X GET http://localhost:8000/api/v1/users/
func GetUsers(w http.ResponseWriter, r *http.Request){
	utils.SendData(w, models.GetUsers() )
}

// curl -X GET http://localhost:8000/api/v1/users/1
func GetUser(w http.ResponseWriter, r *http.Request){
	if user, err := getUserByRequest(r); err != nil{
		utils.SendNotFound(w)
	}else{
		utils.SendData(w, user)
	}
}

//curl -i -H "Content-Type: application/json" -X POST -d '{"username":"New User dos", "password": "password123", "email": "email@email.com"}' http://localhost:8000/api/v1/users/
func CreateUser(w http.ResponseWriter, r *http.Request){
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil{
		utils.SendUnprocessableEntityMessage(w, err.Error())
		return
	}
	
	if err := user.Valid(); err != nil{
		utils.SendUnprocessableEntityMessage(w, err.Error())
		return
	}

	user.SetPassword(user.Password)
	if err := user.Save(); err != nil{
		utils.SendUnprocessableEntityMessage(w, err.Error())
		return	
	}
	utils.SendData(w, user)
}

//curl -i -H "Content-Type: application/json" -X PUT -d '{"username":"Lalo", "password": "change123", "email":"eduardo@codigofacilito.com"}' http://localhost:8000/api/v1/users/1
func UpdateUser(w http.ResponseWriter, r *http.Request){
	user, err := getUserByRequest(r)
	if err != nil{
		utils.SendNotFound(w)
		return
	}

	request := &models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(request); err != nil{
		utils.SendUnprocessableEntityMessage(w, err.Error())
		return
	}
	
	if err := user.Valid(); err != nil{
		utils.SendUnprocessableEntityMessage(w, err.Error())
		return
	}

	user.Username = request.Username
	user.SetPassword(request.Password)
	user.Email = request.Email
	
	if err:= user.Save(); err != nil{
		utils.SendUnprocessableEntityMessage(w, err.Error())
		return
	}
	utils.SendData(w, user)
}

//curl -X DELETE http://localhost:8000/api/v1/users/1 -i
func DeleteUser(w http.ResponseWriter, r *http.Request){
	if user, err := getUserByRequest(r); err != nil{
		utils.SendNotFound(w)
	}else{
		user.Delete()
		utils.SendNoContent(w)
	}
}

func getUserByRequest(r *http.Request) (*models.User, error){
	vars := mux.Vars(r)
	id, _ :=  strconv.Atoi( vars["id"] )
	user := models.GetUserById(id)
	if user.Id == 0{
		return user, errors.New("Usuario inexistente en la base de datos")
	}
	return user, nil
}
