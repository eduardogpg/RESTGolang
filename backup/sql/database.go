package models

// Dcumentation >> https://golang.org/pkg/database/sql/
// The sql package must be used in conjunction with a database driver. 
// https://github.com/golang/go/wiki/SQLDrivers

import (
  "log"
  "../config"
  "database/sql"
  _ "github.com/go-sql-driver/mysql" 
  
)

var db *sql.DB

func init(){
  createConnection()
  createTables()
}

func createConnection(){
  url := config.UrlDatabase()
  if connection, err := sql.Open("mysql", url); err != nil{
    panic(err) //If we do not have access to the database, why continue with the program?
  }else{
    db = connection
  }
}

func CloseConnection(){
  db.Close()
}

func createTables(){
  createTable(UserTable, UserSchema)
}

func createTable(table, schema string){
  if !existsTable(table) {
    execute(schema)
  }else{
    truncateTable(table)
  }
}

func truncateTable(table string){
  execute("TRUNCATE " + table )
}

func existsTable(table string) bool{
  rows, _ := query( "SHOW TABLES LIKE '" + table + "'")
  return rows.Next()
}

func insertData(query string, args ...interface{}) (int64, error){
  if result, err := execute(query, args...); err != nil{
    return int64(0), err
  }else{
    id, _ := result.LastInsertId()
    return id, nil
  }
}

func modifyData(query string, args ...interface{}) (int64, error){
  if result, err := execute(query, args...); err != nil{
    return int64(0), err
  }else{
    rows, _ := result.RowsAffected()
    return rows, nil
  }
}

//Exec executes a query without returning any rows. 
func execute(query string, args ...interface{}) (sql.Result, error){
  result, err := db.Exec(query, args...)
  if err != nil{
    log.Println(err)
  }
  return result, err
}

// Query executes a query that returns rows
func query(query string, args ...interface{}) (*sql.Rows, error) {
  rows, err := db.Query(query, args...)
  if err != nil{
    log.Println(err) //panic
  }
  return rows, err
}








