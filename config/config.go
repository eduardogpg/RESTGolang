package config

import (
  "os"
  "fmt"
  "github.com/eduardogpg/gonv"
)

type Config interface {
  url() string
}

type ServerConfig struct{
  host        string
  port        int
  debug       bool
  templateDir string
}

type DatabaseConfig struct{
  username string
  password string
  host string
  port int
  database string
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
  server.host = gonv.GetStringEnv("HOST", "localhost")
  server.port = gonv.GetIntEnv("PORT", 3000)
  server.debug = gonv.GetBoolEnv("DEBUG", true)
  server.templateDir, _ = os.Getwd() //Obtenemos la dirección actual del servidor
}

func Debug() bool{
  return server.debug
}

func ServerHost() string {
  return server.host;
}

func ServerPort() int {
  return server.port;
}

func UrlDatabase() string {
  return database.url()
}

func UrlServer() string {
  return server.url()
}

func TemplatesDir() string{
  //return fmt.Sprintf("%s/templates/**/*.html", server.templateDir)
  return "templates/**/*.html"
}

func ErrorTemplateDir() string{
  //return fmt.Sprintf("%s/templates/error.html", server.templateDir)
  return "templates/error.html"
}

func AssetsDir() string{
  return fmt.Sprintf("%s/assets", server.templateDir)
}

func (this *ServerConfig) url() string {
  return fmt.Sprintf("%s:%d", this.host, this.port)
}

func (this *DatabaseConfig) url() string{
  return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", this.username, this.password, this.host, this.port, this.database)
}
