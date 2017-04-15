package test

// cd test >> go test

import (
  "../models"

  "strings"
  "testing"
  "fmt"
  "os"
)

var id int

// go test 
// go test -v
// go test -help

const(
  username = "eduardo"
  password = "password123"
  encriptedPassword = "$2a$10$EluB03qsdJESTgAwMZv3AObxAVRDHbCd2Ui4PU.jIv4FrLHrxNtQC"
  email = "eduardo@codigofacilito.com"
)

func TestMain(m *testing.M) { 
  setupFunction()
  result := m.Run()
  teardownFunction()
  os.Exit(result)
}

func setupFunction(){
  createDefaultUser()
}

func teardownFunction(){
}

func createDefaultUser(){
  sql := "INSERT users SET id=?, username=?,password=?,email=?"
  uId, err := models.InsertData(sql, id, username, encriptedPassword, email)
  if err != nil{
    panic(err)
  }
  id = int(uId)
}

//Primera prueba vídeo número 74
func TestDemo(t *testing.T) {
  if str := "Hola mundo!"; str != "Hola mundo!"{
    t.Error("No es posible saludar a los usuarios", nil)
  }
}

func TestNew(t *testing.T){
  user := models.NewUser(username, password, email)
  if user.Username == "" || user.Password == "" || user.Email == "" {
    t.Error("No fue posible crear un nuevo usuario.", nil)
  }
}

func TestPassword(t *testing.T){
  user := models.NewUser(username, password, email)
  if user.Password == password || len(user.Password) != 60 {
    t.Error("No fue posible encriptar el password.", nil)
  }
  t.Log("El nuevo password", user.Password)
}

func TestSave(t *testing.T){
  user := models.NewUser("newUsername", "newPassword", "newEmail")
  if err := user.Save(); err != nil{
    t.Error("No fue posible guardar un nuevo usuario", err)
  }
}

func TestCreate(t *testing.T){
  if _, err := models.CreateUser(username + "1", password, email); err != nil{
    t.Error("No fue posible crear un nuevo usuario.", err)
  }
}

func TestUsername(t *testing.T){
  newUsername := strings.Repeat(username, 10)
  message := fmt.Sprintf("Error 1406: Data too long for column 'username' at row 1")
  
  user := models.NewUser(newUsername, password, email)
  if err := user.Save(); err.Error() != message{
    t.Error("Username demaciado largo.", err)
  }
}

func TestDuplicateUsername(t *testing.T){
  message := fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'username'", username)
  if _, err := models.CreateUser(username, password, email); err.Error() != message{
    t.Error("Username duplicado en la base de datos.", err)
  }
}

func TestGetById(t *testing.T){
  if user := models.GetUser("id", id); user.Id != int64(id){
    t.Error("No es posible obtener el usuario por id.", nil)
  }
}

func TestGetByUsername(t *testing.T){
  if user := models.GetUser("username", username); user.Username != username{
    t.Error("No es posible obtener el usuario por username.", nil)
  }
}

func TestUsers(t *testing.T){
  if users := models.GetUsers(); len(users) == 0{
    t.Error("No es posible obtener los usuarios.", nil)     
  }
}

func TestLogin(t *testing.T){
  if ok := models.LoginUser(username, password); !ok{
    t.Error("No es posible autenticar a un usuario", nil)
  }
}
