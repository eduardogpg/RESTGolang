package handler

import(
	"net/http"
  "fmt"
)

func LoginUser(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Eduardo")
}
