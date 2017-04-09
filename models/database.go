package models

// Dcumentation >> https://golang.org/pkg/database/sql/
// The sql package must be used in conjunction with a database driver. 
// https://github.com/golang/go/wiki/SQLDrivers

import (
  _ "github.com/go-sql-driver/mysql" 
  "database/sql"
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
  rows := query( "SHOW TABLES LIKE '" + table + "'")
  return rows.Next()
}

// https://golang.org/pkg/database/sql/#DB.Exec
func insertData(query string, args ...interface{}) int64{
  result := execute(query, args...)
  id, _ := result.LastInsertId()
  return id
}

func modifyData(query string, args ...interface{}) int64{
  res := execute(query, args...)
  rows, _ := res.RowsAffected()
  return rows
}


//Exec executes a query without returning any rows. 
func execute(query string, args ...interface{}) sql.Result {
  if result, err := db.Exec(query, args...); err != nil{
    panic(err)
  }else{
    return result
  }
}

// Query executes a query that returns rows
func query(query string, args ...interface{}) *sql.Rows {
  if rows, err := db.Query(query, args...); err != nil{
    panic(err)
  }else{
    return rows
  }
}








