package models

import (
  "log"
  "../config"
  
  "database/sql"
  _ "github.com/go-sql-driver/mysql" 
)

var db *sql.DB
var debug bool

func init(){
  CreateConnection()
  debug = config.Debug()
}

func CreateConnection(){
  if Connection() != nil{
    return
  }

  if connection, err := sql.Open("mysql", config.UrlDatabase()); err != nil{
    panic(err) 
  }else{
    db = connection
  }
} 

func Ping(){
  if err := db.Ping(); err != nil{
    log.Println(err)
    panic(err)
  }
}

func Connection() *sql.DB{
  return db
}

func CloseConnection(){
  db.Close()
}

func CreateTables(){
  createTable(UserTable, UserSchema)
}

func createTable(table, schema string){
  if !existsTable(table) {
    Execute(schema)
  }else{
    truncateTable(table)
  }
}

func existsTable(table string) bool{
  rows, _ := Query( "SHOW TABLES LIKE '" + table + "'")
  return rows.Next()
}

func truncateTable(table string){
  Execute("TRUNCATE " + table )
}

func InsertData(query string, args ...interface{}) (int64, error){
  if result, err := Execute(query, args...); err != nil{
    return int64(0), err
  }else{
    id, _ := result.LastInsertId()
    return id, nil
  }
}

func ModifyData(query string, args ...interface{}) (int64, error){
  if result, err := Execute(query, args...); err != nil{
    return int64(0), err
  }else{
    rows, _ := result.RowsAffected()
    return rows, nil
  }
}

//Exec executes a query without returning any rows. 
func Execute(query string, args ...interface{}) (sql.Result, error){
  result, err := db.Exec(query, args...)
  if err != nil && !debug { 
    log.Println(err)
  }
  return result, err
}

// Query executes a query that returns rows
func Query(query string, args ...interface{}) (*sql.Rows, error) {
  rows, err := db.Query(query, args...)
  if err != nil && !debug{ 
    log.Println(err)
  }
  return rows, err
}

