package config

import (
  "log"

  "github.com/caarlos0/env"
  "github.com/jinzhu/gorm"

  _ "github.com/go-sql-driver/mysql"
)

type DataBaseConfig struct {
  Engine    string    `env:"engine" envDefault:"mysql"`
  Username  string    `env:"dbuser" envDefault:"root"`
  Password  string    `env:"dbpassword" envDefault:""`
  Host      string    `env:"dbhost" envDefault:"localhost"`
  Port      int       `env:"dbport" envDefault:"3306"`
  Database  string    `env:"database" envDefault:"go_web"`
}

var dataConfig DataBaseConfig //privado
var Connection *gorm.DB

func (this *DataBaseConfig) getUrl() string{
  return this.Username + ":" + this.Password + "@/" + this.Database
}

func databaseInitializer(){
  if env.Parse(&dataConfig) != nil {
    log.Println("No fue posible obtener las configuraciones de la base de datos!")
  }else{
    createGormConnection()
  }
}

func createGormConnection(){
  db, err := gorm.Open(dataConfig.Engine, dataConfig.getUrl() )
  if err != nil{
    log.Println("No fue posible realizar la conexión a la base de datos!")
    return
  }
  Connection = db
}

func closeGormConnection(){
  Connection.Close()
}


