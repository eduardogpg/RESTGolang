package test

import(
  "fmt"
  "time"
  "testing"
  "math/rand"
  "../models"
)

var user *models.User

const(
  id = 1
  username = "eduardo_gpg"
  password =  "password"
  encryptedPassword = "$2a$10$4BxwX4tvOY2UifHsrp8eBOHzcp0IjPkSvd/iM3VqOh/h1mrJrp9Vi"
  email = "eduardo@codigofacilito.com"
  invalidEmail = "eduardo"
  createdDate = "2017-08-17"
)

func TestNewUser(t *testing.T){
  user := models.NewUser(username,password, email) //El pass ya encriptado no coincide
  if !objectEquals(user) { //t.Log("El password encriptad es ", user.Password)
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
    t.Error("No es posible insertar el objeto", nil)
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
  message := fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'username'", username)
  if err.Error() != message {
    t.Error("Username duplicado en la base de datos.", err)
  }
}

func TestPassword(t *testing.T){
  user := models.NewUser(username, password, email)
  if user.Password == password{
    t.Error("No fue posible cifrar el password")
  }
}

func TestPasswordLength(t *testing.T){
  user := models.NewUser(username, password, email)
  if len(user.Password) != 60{
    t.Error("No fue posible cifrar el password") 
  }
}

func TestValidLogin(t *testing.T){
  if valid := models.Login(username, password); !valid{
    t.Error("No es posible realizar la autenticación")
  }
}

func TestValidEmail(t *testing.T){
  if err := models.ValidEmail(email); err != nil{
    t.Error("No es posible validar el Email")
  }
}

func TestInValidEmail(t *testing.T){
  err := models.ValidEmail(invalidEmail)
  if err == nil || err.Error() != "Formato invalido de Email"{
    t.Error("Es posible registrar un Email invalido")
  }
}

func TestGet(t *testing.T){
  user := models.GetUser("id", id)
  if !objectEquals(user) || !createDateEquals(user.CreatedDate) { //t.Log(user.CreatedDate)
    t.Error("No es posible obtener el usuario")
  }
}

func TestGetUsers(t *testing.T){
  users := models.GetUsers()
  if len(users) == 0{
    t.Error("No es posible obtener los usuarios")
  }
}
/*
func TestDeleteUser(t *testing.T){
  if err := user.Delete(); err != nil{
    t.Error("No es posible eliminar el usuario")
  }
}
*/

/*Funciones internas*/
func random_username() string{
  return fmt.Sprintf("%s/%d", username, rand.Intn(1000))
}

func createDateEquals(date time.Time) bool{
  t, _ := time.Parse("2006-01-02", createdDate)
  return date == t
}

func objectEquals(user *models.User) bool{
  // return user.Username == username && user.Password == password && user.Email == email
  return user.Username == username && user.Email == email
}


