package utils

import(
  "fmt"
  "sync"
  "time"
  "net/http"
  "../models"
  "github.com/satori/go.uuid"
)

const(
  cookieName = "go_session"
  cookieExprires = 24 * 2 * time.Hour
)

// https://golang.org/pkg/sync/#RWMutex
var Sessions = struct{
  m map[string] *models.User
  sync.RWMutex
}{m: make(map[string] *models.User)}

//https://golang.org/src/net/http/cookie.go
func SetSession(user *models.User, w http.ResponseWriter, r *http.Request){
  Sessions.RLock()
  defer Sessions.RUnlock()

  uuid := uuid.NewV4().String()
  Sessions.m[uuid] = user

  cookie := &http.Cookie{
    Name: cookieName,
    Value: uuid,
    Expires: time.Now().Add(cookieExprires),
    Path: "/", //The cookie will be available to all pages and subdirectories.
  }
  http.SetCookie(w, cookie)
}

func getUserBySession(key string) (*models.User){
  Sessions.RLock()
  defer Sessions.RUnlock()

  if user, ok := Sessions.m[key]; ok{
    return user
  }
  return nil
}

func GetUser(r *http.Request) (*models.User) {
  val := getValCookie(r)
  return getUserBySession(val)
}

func getValCookie(r *http.Request) string {
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
  return getValCookie(r) != ""
}
