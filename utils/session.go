package utils

import(
  "fmt"
  "net/http"
)
//http://localhost:3000/login

const sessionName = "sisisisisis"
//https://mschoebel.info/2014/03/09/snippet-golang-webapp-login-logout/

//https://golang.org/src/net/http/cookie.go
//https://astaxie.gitbooks.io/build-web-application-with-golang/en/06.2.html
func SetSession(w http.ResponseWriter, r *http.Request, userName string){
  _, err := r.Cookie(sessionName) //Nos regresa un cookie como &
  if err != nil{
    http.SetCookie(w, &http.Cookie{
      Name: sessionName,
      Value: userName,
      Path: r.URL.Path,
    })

    x, y := r.Cookie(sessionName); 
    fmt.Println(x)
    fmt.Println(x.Path)
    fmt.Println(y)    

  }
}

func UpdateSession(w http.ResponseWriter, r *http.Request){
  if cookie, err := r.Cookie(sessionName); err == nil {
    cookie.Path = r.URL.Path
    http.SetCookie(w, cookie)
  }
}

func DeleteSession(w http.ResponseWriter){
  cookie := &http.Cookie{
    Name:   sessionName,
    Value:  "",
    Path: "/logout",
    MaxAge: -1,
  }
  http.SetCookie(w, cookie)
}

func HasSession(r *http.Request) bool{
  cookie, err := r.Cookie(sessionName); 
  fmt.Println(cookie)
  fmt.Println(err)
  return err == nil && cookie != nil && cookie.Value != ""
}

