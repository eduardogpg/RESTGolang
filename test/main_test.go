package test

import(
  "testing"
  "os"
  "fmt"
  "../models"
)

// https://golang.org/pkg/testing/#hdr-Main
func TestMain(m *testing.M) {
  beforeTest()
  result := m.Run() //Run runs the tests. It returns an exit code to pass to os.Exit.
  afterTest()
  os.Exit(result) //result is a int
}

func beforeTest(){
  fmt.Println(">> Before the suite")
  models.CreateConnection()
  models.CreateTables()
  createDefaultUser()
}

func afterTest(){
  fmt.Println(">> After the suite")
  models.CloseConnection()
}


func createDefaultUser(){
  sql := fmt.Sprintf("INSERT users SET id='%d', username='%s', password='%s', email='%s', created_date='%s' ",id, username, encryptedPassword, email, createdDate)
  _, err := models.Execute(sql)
  if err != nil{
    panic(err)
  }
  user = &models.User{Id: int64(id), Username:username, Password:password, Email:email}
}




