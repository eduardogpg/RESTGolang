package models

// Documentation
// https://golang.org/pkg/database/sql/

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

func CreateConnection(){
  url := getDatabaseURL()
  if db, err := sql.Open(engineSQL, url); err != nil{
    log.Fatal(err)
  }else{
    connection = db
    createTables()
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
  if existsTable(table) == false{
    if _, err := connection.Exec(schema); err != nil{
      panic(err)
    }
  }
}

func existsTable(table string) bool{
  if rows, err := connection.Query("SHOW TABLES LIKE '"+ table +"' "); err != nil{
    panic(err)
  }else{
    return rows.Next()
  }
}

// https://golang.org/pkg/database/sql/#DB.Exec
func insertData(query string, args ...interface{}) int64{
  result := executeSql(query, args...)
  id, _ := result.LastInsertId()
  return id
}

func executeSql(query string, args ...interface{}) sql.Result {
  if result, err := connection.Exec(query, args...); err != nil{
    panic(err)
  }else{
    return result
  }
}











