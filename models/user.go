package models

import(
  "fmt"
  "time"
  "regexp"
  "errors"
  "golang.org/x/crypto/bcrypt"
)

const UserTable string = "users"

type User struct{
  Id            int64     `json:"id"`
  Username      string    `json:"username"`
  Password      string    `json:"password"`
  Email         string    `json:"email"`
  CreatedDate   time.Time
}

type Users []User

const UserSchema string = `CREATE TABLE users (
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(30) NOT NULL UNIQUE,
        password VARCHAR(60) NOT NULL,
        email VARCHAR(50),
        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (this *User) Save() error {
  if this.Id > 0{
    return this.Update()
  }else{
    return this.Insert()
  }
}

func (this *User) Delete() error {
  sql := "DELETE FROM users WHERE id=?"
  _, err := ModifyData(sql, 2)
  return err
}

func (this *User) Update() error {
  sql := "UPDATE users SET username=?,password=?,email=? WHERE id=?"
  _, err := ModifyData(sql, this.Username, this.Password, this.Email, this.Id)
  return err
}

func (this *User) Insert() error {
  sql := "INSERT users SET username=?,password=?,email=?"
  id, err := InsertData(sql, this.Username, this.Password, this.Email)
  this.Id = id
  return err
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/en/09.5.html
func (this *User) SetPassword(password string){
  hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  this.Password = string(hash)
}

func NewUser(username, password, email string) *User {
  user :=&User{Username: username, Email: email }
  user.SetPassword(password)
  return user 
}

func ValidEmail(email string) error{
  if !emailRegexp.MatchString(email) {
    return errors.New("Formato invalido de Email")
  }
  return nil
} 

func Login(username, password string) bool{
  user := GetUser("username", username)
  err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
  return err == nil
}

func CreateUser(username, password, email string) (*User, error){
  user := NewUser(username, password, email)
  return user, user.Save()
}

func GetUser(field string, conditional interface{}) *User{
  sql := fmt.Sprintf("SELECT id, username, password, email, created_date FROM users WHERE %s=?", field)
  user := &User{}
  row, err := Query(sql, conditional);
  
  if err != nil{
    return user
  }

  for row.Next() {
    row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedDate)
  }
  return user
}

func GetUsers() Users{
  sql := "SELECT id, username, password, email, created_date FROM users"
  row, err := Query(sql)
  users := Users{}

  if err != nil{
    return users
  }

  for row.Next() {
    user := User{}
    row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedDate)
    users = append(users, user)
  }
  return users
}

func LoginUser(username, password string) bool {
  user := GetUser("username", username)
  err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password));
  return err == nil
}
