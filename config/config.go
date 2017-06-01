package config

import (
  "github.com/eduardogpg/gonv"
  "fmt"
)

type ServerConfig struct{
  port  int
  debug bool
}

type DatabaseConfig struct{
  username string
  password string
  host string
  port int
  database string
  debug bool
}

var server *ServerConfig
var database *DatabaseConfig

func init(){
  database = &DatabaseConfig{}
  database.username = gonv.GetStringEnv("USERNAME", "root")
  database.password = gonv.GetStringEnv("PASSWORD", "")
  database.host = gonv.GetStringEnv("HOST", "localhost")
  database.port = gonv.GetIntEnv("PORT", 3306)
  database.database = gonv.GetStringEnv("DATABASE", "project_go_web")
  
  server = &ServerConfig{}
  server.port = gonv.GetIntEnv("PORT", 3000)
  server.debug = gonv.GetBoolEnv("DEBUG", true)
}

func GetDebug() bool{
  return server.debug
}

func GetUrlDatabase() string {
  return database.url()
}

func (this *DatabaseConfig) url() string{
  return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", this.username, this.password, this.host, this.port, this.database)
}