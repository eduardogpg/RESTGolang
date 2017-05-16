package test

import(
  "testing"
)

func TestHolaMundo(t *testing.T){
  str := "hola mundo"
  
  if str != "hola mundo"{ 
    t.Error("No es posible saludar a los usuarios", nil)
  }
}