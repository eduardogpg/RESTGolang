package utils

import(
  "fmt"
  "sync"
  "time"
  "net/http"
  "github.com/satori/go.uuid"
)

const(
  cookieName = "go_session"
  cookieExprires = 24 * 2 * time.Hour
)


//https://golang.org/src/net/http/cookie.go
func SetSession(w http.ResponseWriter, r *http.Request){
  cookie := &http.Cookie{
    Name: cookieName,
    Value: uuid.NewV4().String(), 
    Expires: time.Now().Add(cookieExprires),
    Path: "/", //The cookie will be available to all pages and subdirectories.
  }
  http.SetCookie(w, cookie)
}

func GetValCookie(r *http.Request) string {
  if cookie, err := r.Cookie(cookieName); err == nil {
    val := cookie.Value
    fmt.Println(val)
    return val
  }else{
    fmt.Println(err)  
  }
  return ""
}

func DeleteSession(w http.ResponseWriter) {
  cookie := &http.Cookie{
    Name: cookieName,
    Value: "",
    MaxAge: -1,
    Path: "/",
  }
  http.SetCookie(w, cookie)
}

func IsAuthenticated(r *http.Request) bool{
  return GetValCookie(r) != ""
}
