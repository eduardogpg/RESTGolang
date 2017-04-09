package test

// cd test >> go test

import (
  "testing"
  "../models"
  "os"
)

func TestMain(m *testing.M) { 
  setup()
  retCode := m.Run()
  teardown()
  os.Exit(retCode)
}

func setup(){
  models.InitializeDatabase()
  createUsers()
}

func teardown(){
  models.CloseConnection()
}

func createUsers(){
  models.CreateUser("eduardo", "password", "eduardo@codigofacilito.com")
}

func TestNewUser(t *testing.T) {
  user := models.NewUser("test", "test", "test@testing.com")
  if user.Username != "test"{
    t.Error("No es posible crear el usuario.", nil)
  }
}

func TestCreateUser(t *testing.T){
  user := models.CreateUser("test", "test", "test@testing.com")
  if err := user.GetErrors(); err != nil {
    t.Error("No es posible persistir el usuario", err)
  }
}

func TestUniqueUser(t *testing.T){
  user := models.CreateUser("eduardo", "test", "test@testing.com")
  if err := user.GetErrors(); err != nil {
    t.Error(nil, err)
  }
}


/*
  1.- Establecer password
  2.- Unique
  3.- 
*/


