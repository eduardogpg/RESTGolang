package handler

import(
  "net/http"
  "strconv"
  
  "../models"
  "../utils"

  "encoding/json"
  "github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
  utils.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request){
  if user, err := GetUserByRequest(r); err != nil{ 
    utils.SendNotFound(w)
  }else{
    utils.SendData(w, user)
  }
}

func CreateUser(w http.ResponseWriter, r *http.Request){
  user := models.User{}
  decoder := json.NewDecoder(r.Body)

  if err := decoder.Decode(&user); err != nil{
    utils.SendUnprocessableEntity(w)
  }else{
    user = models.SaveUser(user)
    utils.SendData(w, user) 
  }
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
  user, err := GetUserByRequest(r)
  
  if err != nil{ 
    utils.SendBadRequest(w)
    return
  }   
  
  userR := models.User{}
  decoder := json.NewDecoder(r.Body)

  if err := decoder.Decode(&userR); err != nil{
    utils.SendBadRequest(w)
    return
  }
  
  user = models.UpdateUser(user, userR.Username, userR.Password)
  utils.SendData(w, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
  if user, err := GetUserByRequest(r); err != nil{ 
    utils.SendBadRequest(w)
  }else{
    models.DeleteUser(user.Id)
    utils.SendNoContent(w)
  }
}

func GetUserByRequest(r *http.Request)(models.User, error){
  vars := mux.Vars(r)
  id, _ := strconv.Atoi(vars["id"]) 

  if user, err := models.GetUser(id); err != nil{ 
    return user, err
  }else{
    return user, nil
  }
}
