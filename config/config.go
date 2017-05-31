package config

import (
  "fmt"
  "github.com/caarlos0/env"
)

type configInterface interface {
    Url() string
}

type DatabaseConfig struct {
  Username    string    `env:"USERNAME" envDefault:"root"`
  Password    string    `env:"PASSWORD" envDefault:""`
  Database    string    `env:"DATABASE" envDefault:"REST_GOLANG"`
  Port        int       `env:"PORT" envDefault:"3306"`
  Host        string    `env:"HOST" envDefault:"localhost"`
  Production  bool      `env:"PRODUCTION" envDefault:"false"`
}

var database *DatabaseConfig

func init() {
  database = &DatabaseConfig{}
  if err := env.Parse(database); err != nil {
    panic(err)
  }
}

//"<username>:<pw>@tcp(<HOST>:<port>)/<dbname>"
func (this *DatabaseConfig) Url() string{
  return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",this.Username, this.Password, this.Host, this.Port, this.Database)
}

func UrlDatabase() string{
  return database.Url()
}

func DebugDatabase() bool{
  return !database.Production
}


