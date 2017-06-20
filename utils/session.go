package utils

import(
  "fmt"
  "time"
  "net/http"
  "github.com/satori/go.uuid"
)

const(
  cookieName = "rosadito"
  cookieExprires = 24 * 2 * time.Hour
)

//https://golang.org/src/net/http/cookie.go
func SetSession(w http.ResponseWriter, r *http.Request){
  _, err := r.Cookie(cookieName) //https://golang.org/pkg/net/http/#Request.Cookie
  if err != nil{
    cookie := &http.Cookie{
      Name: cookieName,
      Value: uuid.NewV4().String(), 
      Expires: time.Now().Add(cookieExprires),
    }
    http.SetCookie(w, cookie)
  }
}

func GetValCookie(r *http.Request) string {
  if cookie, err := r.Cookie(cookieName); err == nil {
    val := cookie.Value
    return val
  }else{
    fmt.Println(err)  
  }
   
  return ""
}

func DeleteSession(response http.ResponseWriter) {
  cookie := &http.Cookie{
    Name: cookieName,
    Value: "",
    MaxAge: -1,
  }
  http.SetCookie(response, cookie)
}

func IsAuthenticated(r *http.Request) bool{
  return GetValCookie(r) != ""
}
