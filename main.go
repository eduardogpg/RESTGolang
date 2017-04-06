package main

import(
  "log"
  "fmt"
  
	"./app"
)

func main() {
  
  server := app.CreateServer()

  fmt.Println("Servidor a la escucha en el puerto 3000" )
	log.Fatal(server.ListenAndServe())
}
