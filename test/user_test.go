package test

import(
  "fmt"
  "testing"
  "math/rand"
  "../models"
)

var user *models.User

const(
  id = 1
  username = "eduardo_gpg"
  password = "password"
  email = "eduardo@codigofacilito.com"
  message_duplicate_username = "Error 1062: Duplicate entry 'eduardo_gpg' for key 'username'"
)

func TestNewUser(t *testing.T){
  user := models.NewUser(username,password, email)
  if user.Username != username || user.Password !=  password || user.Email != email {
    t.Error("No es posible crear el objeto")
  }
}

func TestSave(t *testing.T){
  user := models.NewUser(random_username(), password, email)
  if err := user.Save(); err != nil{
    t.Error("No es posible crear el usuario", err)
  }
}

func TestCreateUser(t *testing.T){
  _, err := models.CreateUser(random_username(), password, email)
  if err != nil{
    t.Error("No es posible insertar el objeto", nil) //aqui coloque err
  }
}

func TestUniqueUsername(t *testing.T){
  _, err := models.CreateUser(username, password, email)
  if err == nil{
    t.Error("Es posible insertar un username duplicados")
  }
}

func TestDuplicateUsername(t *testing.T){
  _, err := models.CreateUser(username, password, email);
  if err.Error() != message_duplicate_username {
    t.Error("Username duplicado en la base de datos.", err)
  }
}

func TestGet(t *testing.T){
  user := models.GetUser("id", id)
  if user.Username != username || user.Password  != password || user.Email != email{
    t.Error("No es posible obtener el usuario")
  }
}

func TestGetUsers(t *testing.T){
  users := models.GetUsers()
  if len(users) == 0{
    t.Error("No es posible obtener los usuarios")
  }
}

func TestUpdateUser(t *testing.T){
  user.Email = "cambio@email.com"
  if err := user.Save(); err != nil{
    t.Error("No es posible actualizar el usuario")
  }

  if user.Email == email{
    t.Error("No es posible actualizar el usuario") 
  }
}

func TestDeleteUser(t *testing.T){
  if err := user.Delete(); err != nil{
    t.Error("No es posible eliminar el usuario")
  }
}

func random_username() string{
  return fmt.Sprintf("%s/%d", username, rand.Intn(1000))
}



