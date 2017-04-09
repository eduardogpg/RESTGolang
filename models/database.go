package models

// Dcumentation >> https://golang.org/pkg/database/sql/
// The sql package must be used in conjunction with a database driver. 
// https://github.com/golang/go/wiki/SQLDrivers

import (
  _ "github.com/go-sql-driver/mysql" 
  "database/sql"
  "log"
)

var connection *sql.DB
const engineSQL string = "mysql"

const username string = "root"
const password string = ""
const database string = "REST_GOLANG"

func InitializeDatabase(){
  createConnection()
  createTables()
}

func createConnection(){
  url := getDatabaseURL()
  if db, err := sql.Open(engineSQL, url); err != nil{
    log.Fatal(err)
  }else{
    connection = db
  }
}

func CloseConnection(){
  connection.Close()
}

func getDatabaseURL() string{
  return username + ":" + password + "@/" + database + "?charset=utf8"
}

func createTables(){
  createTable(UserTable, UserSchema)
}

func createTable(table, schema string){
  if !existsTable(table) {
    if _, err := connection.Exec(schema); err != nil{
      panic(err)
    }
  }
}

func existsTable(table string) bool{
  rows := executeQuery( "SHOW TABLES LIKE " + table )
  return rows.Next()
}

// https://golang.org/pkg/database/sql/#DB.Exec
func insertData(query string, args ...interface{}) int64{
  result := executeSql(query, args...)
  id, _ := result.LastInsertId()
  return id
}

func insertData(query string, args ...interface{}) int64{
  result := executeSql(query, args...)
  id, _ := result.LastInsertId()
  return id
}

//Exec executes a query without returning any rows. 
func executeSql(query string, args ...interface{}) sql.Result {
  if result, err := connection.Exec(query, args...); err != nil{
    panic(err)
  }else{
    return result
  }
}

// Query executes a query that returns rows
func executeQuery(query string, args ...interface{}) *sql.Rows {
  if result, err := connection.Query(query, args...); err != nil{
    panic(err)
  }else{
    return result
  }
}








