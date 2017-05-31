package handlers

import(
	"log"
	"strconv"
	"errors"
	"net/http"
	"../models"
	"encoding/json"
	"github.com/gorilla/mux"
)

// curl -X GET http://localhost:8000/api/v1/users/
func GetUsers(w http.ResponseWriter, r *http.Request){
	SendData(w, models.GetUsers() )
}

// curl -X GET http://localhost:8000/api/v1/users/1
func GetUser(w http.ResponseWriter, r *http.Request){
	if user, err := getUserByRequest(r); err != nil{
		SendNotFound(w)
	}else{
		SendData(w, user)
	}
}

//curl -i -H "Content-Type: application/json" -X POST -d '{"username":"New User dos", "password": "password123", "email": "email@email.com"}' http://localhost:8000/api/v1/users/
func CreateUser(w http.ResponseWriter, r *http.Request){
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil{
		SendUnprocessableEntity(w)
	}
	
	user.SetPassword(user.Password)
	if err := user.Save(); err != nil{
		SendUnprocessableEntity(w)	
	}
	SendData(w, user)
}

//curl -i -H "Content-Type: application/json" -X PUT -d '{"username":"Lalo", "password": "change123", "email":"eduardo@codigofacilito.com"}' http://localhost:8000/api/v1/users/1
func UpdateUser(w http.ResponseWriter, r *http.Request){
	user, err := getUserByRequest(r)
	if err != nil{
		SendNotFound(w)
		return
	}
	
	request := &models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(request); err != nil{
		log.Println("Primer if")
		SendUnprocessableEntity(w)
		return
	}

	user.Username = request.Username
	user.SetPassword(request.Password)
	user.Email = request.Email
	
	if err:= user.Save(); err != nil{
		log.Println("Segundo if")
		SendUnprocessableEntity(w)
	}
	SendData(w, user)
}

//curl -X DELETE http://localhost:8000/api/v1/users/1 -i
func DeleteUser(w http.ResponseWriter, r *http.Request){
	if user, err := getUserByRequest(r); err != nil{
		SendNotFound(w)
	}else{
		user.Delete()
		SendNoContent(w)
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