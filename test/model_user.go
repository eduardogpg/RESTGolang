package test

import (
  "testing"
  "../models"
  "os"
)

func TestMain(m *testing.M) { 
  setup()
  retCode := m.Run()
  teardown()
  os.Exit(retCode)
}

func setUp(){

}

func teardown(){
  
}