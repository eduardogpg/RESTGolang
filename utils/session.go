package utils

import(
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
  sync.RWMutex        //RWMutex is a reader/writer mutual exclusion lock.
}{m: make(map[string] *models.User)}

//https://golang.org/src/net/http/cookie.go
func SetSession(user *models.User, w http.ResponseWriter, r *http.Request){
  Sessions.Lock()
  defer Sessions.Unlock()
  
  uuid := uuid.NewV4().String()
  Sessions.m[uuid] = user

  cookie := &http.Cookie{
    Name: cookieName,
    Value: uuid,
    Expires: time.Now().Add(cookieExprires),
    Path: "/",
  }
  http.SetCookie(w, cookie)
}

func getUserBySession(key string) (*models.User){
  Sessions.Lock()
  defer Sessions.Unlock()

  if user, ok := Sessions.m[key]; ok{
    return user
  }
  return &models.User{}
}

func GetUser(r *http.Request) (*models.User) {
  val := GetValCookie(r)
  return getUserBySession(val)
}

func GetValCookie(r *http.Request) string {
  if cookie, err := r.Cookie(cookieName); err == nil {
    return cookie.Value
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
