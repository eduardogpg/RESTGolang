package models

import (
  "../config"
  
  "github.com/jinzhu/gorm"
   _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func init(){
  createConnection()
  createTables()
}

func createConnection() {
  url := config.UrlDatabase()
  if connection, err := gorm.Open("mysql", url); err != nil{
    panic(err) 
  }else{
    db = connection
  }
}

func CloseConnection(){
  db.Close()
}
  
func createTables(){
  db.CreateTable(&User{})
}

func dropTables(){
  db.DropTable(&User{})
}

