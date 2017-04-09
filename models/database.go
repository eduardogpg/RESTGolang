package models

// Dcumentation >> https://golang.org/pkg/database/sql/
// The sql package must be used in conjunction with a database driver. 
// https://github.com/golang/go/wiki/SQLDrivers

import (
  _ "github.com/go-sql-driver/mysql" 
  "database/sql"
  "log"
)

var db *sql.DB
const DataEngine string = "mysql"

const Username string = "root"
const Password string = ""
const DataName string = "REST_GOLANG"

func InitializeDatabase(){
  createConnection()
  createTables()
}

func createConnection(){
  url := GetDatabaseURL()
  if connection, err := sql.Open(DataEngine, url); err != nil{
    panic(err)
  }else{
    db = connection
  }
}

func CloseConnection(){
  db.Close()
}

func GetDatabaseURL() string{
  return Username + ":" + Password + "@/" + DataName + "?charset=utf8"
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

// https://golang.org/pkg/database/sql/#DB.Exec
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
    log.Println(err)
  }
  return rows, err
}








