package main

import(
  "fmt"
  "./models"
  )

func init(){
  
}

func main() {
  user := models.NewUser("eduardo", "password", "eduardo@codigofacilito.com")
  user.Save()
  fmt.Println("Usuario almacenado correctamente!")

  user.Username = "eduardo_gpg"
  user.SetPassword("password123")
  user.Email = "eduardo_gpg@codigofacilito.com"
  user.Save()
  fmt.Println("Usuario actualizado correctamente!")

  user = models.GetUser("id", 1)
  fmt.Println(user)

  user.Delete()
  
  users := models.GetUsers()
  fmt.Println(users)

}