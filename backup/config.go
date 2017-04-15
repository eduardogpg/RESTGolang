package config

import (
  "fmt"
  _ "github.com/caarlos0/env"
  "../utils"
)

type configInterface interface {
    getUrl() string
}

type DatabaseConfig struct {
  Username    string
  Password    string
  Database    string
  Port        int
}

type ServerConfig struct {
  Host         string
  Port         int
  IsProduction bool
}

var database *DatabaseConfig
var server *ServerConfig

func (this *DatabaseConfig) getUrl() string{
  return fmt.Sprintf("%s:%s@/%s?charset=utf8", this.Username, this.Password, this.Database)
}

func (this *ServerConfig) getUrl() string{
  return fmt.Sprintf("%s:%d", this.Host, this.Port)
}


func init() {
  server = &ServerConfig{}
  database = &DatabaseConfig{}
  
  server.Host = utils.GetStringEnv("HOST", "localhost")
  server.Port = utils.GetIntEnv("PORT", 8000)
  server.IsProduction = utils.GetBoolEnv("PRODUCTION", false)

  database.Username = utils.GetStringEnv("USERNAME", "root")
  database.Password = utils.GetStringEnv("PASSWORD", "")
  database.Database = utils.GetStringEnv("DATABASE", "REST_GOLANG")
  database.Port = utils.GetIntEnv("PORT", 3306)

}

func UrlDatabase() string{
  return database.getUrl()
}

func UrlServer() string{
  return server.getUrl()
}

