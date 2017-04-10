package config

import (
  "fmt"
  "github.com/caarlos0/env"
)

type configInterface interface {
    getUrl() string
}

type DatabaseConfig struct {
  Username    string    `env:"USERNAME" envDefault:"root"`
  Password    string    `env:"PASSWORD" envDefault:""`
  Database    string    `env:"DATABASE" envDefault:"REST_GOLANG"`
  Port        int       `env:"PORT" envDefault:"3306"`
}

type ServerConfig struct {
  Host         string   `env:"HOST" envDefault:"localhost"`
  Port         int      `env:"PORT" envDefault:"3000"`
  IsProduction bool     `env:"PRODUCTION" envDefault:"false"`
}

var database *DatabaseConfig
var server *ServerConfig

func (this *DatabaseConfig) getUrl() string{
  return fmt.Sprintf("%s:%s@/%s?charset=utf8", this.Username, this.Password, this.Database)
}

func (this *ServerConfig) getUrl() string{
  return fmt.Sprintf("%s:%d", this.Host, this.Port)
}

//http://stackoverflow.com/questions/9059129/gae-go-init-call-it-multiple-times
func init() {
  server = &ServerConfig{}
  database = &DatabaseConfig{}
  
  if err := env.Parse(server); err != nil {
    panic(err)
  }

  if err := env.Parse(database); err != nil {
    panic(err)
  }
}


func UrlDatabase() string{
  return database.getUrl()
}

func UrlServer() string{
  return server.getUrl()
}

