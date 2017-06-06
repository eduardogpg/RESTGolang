package controllers

import(
  "time"
  "net/http"
  "github.com/satori/go.uuid"
)

const(
  cookieName = "SessionID"
  //cookieExprires = 24 * 2 * time.Hour
  cookieExprires = 1  * time.Minute
)

func setSession(w http.ResponseWriter, r *http.Request){
  _, err := r.Cookie(cookieName)
  if err != nil{
    cookie := &http.Cookie{
      Name: cookieName,
      Value: uuid.NewV4().String(), 
      Expires: time.Now().Add(cookieExprires),
    }
    http.SetCookie(w, cookie)
  }
  //5b a2c
}

func getValCookie(r *http.Request) string {
  if cookie, err := r.Cookie(cookieName); err == nil {
    val := cookie.Value
    return val
  }
  return ""
}

func deleteSession(response http.ResponseWriter) {
  cookie := &http.Cookie{
    Name: cookieName,
    Value: "",
    MaxAge: -1,
  }
  http.SetCookie(response, cookie)
}

func isAuthenticated(r *http.Request) bool{
  return getValCookie(r) != ""
}