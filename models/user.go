package models


import(
  "golang.org/x/crypto/bcrypt"
  "fmt"
)

const UserTable string = "users"

type User struct{
  Id        int64   `json:"id"`
  Username  string  `json:"username"`
  Password  string  `json:"password"`
  Email     string  `json:"email"`
}

type Users []User

const UserSchema string = `CREATE TABLE users (
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(30) NOT NULL UNIQUE,
        password VARCHAR(60) NOT NULL,
        email VARCHAR(50),
        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`


func (this *User) Save() error {
  if this.Id > 0{
    return this.update()
  }else{
    return this.insert()
  }
}

func (this *User) Delete() error {
  sql := "DELETE FROM users WHERE id=?"
  _, err := modifyData(sql, this.Id)
  return err
}

func (this *User) update() error {
  sql := "UPDATE users SET username=?,password=?,email=? WHERE id=?"
  _, err := modifyData(sql, this.Username, this.Password, this.Email, this.Id)
  return err
}

func (this *User) insert() error {
  sql := "INSERT users SET username=?,password=?,email=?"
  id, err := InsertData(sql, this.Username, this.Password, this.Email)
  this.Id = id
  return err
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/en/09.5.html
// http://stackoverflow.com/questions/7465204/maximum-mysql-user-password-length 
// 32 caracteres como maximo
func (this *User) SetPassword(password string){
  hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  this.Password = string(hash) //60 caracteres
}


func NewUser(username, password, email string) *User{
  user :=&User{Username: username, Email: email }
  user.SetPassword(password)
  return user 
}

func CreateUser(username, password, email string) (*User, error){
  user := NewUser(username, password, email)
  return user, user.Save()
}


/*
func GetUser(id int) User{
  sql := "SELECT id, username, password, email FROM users WHERE id=?"
  user := User{}
  row, _ := query(sql, conditional);

  for row.Next() {
    row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
  }
  return user
}
*/

func GetUser(field string, conditional interface{}) *User{
  sql := fmt.Sprintf("SELECT id, username, password, email FROM users WHERE %s=?", field)
  user := &User{}
  row, _ := query(sql, conditional);

  for row.Next() {
    row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
  }
  return user
}

func GetUsers() Users{
  sql := "SELECT id, username, password, email FROM users"
  row, _ := query(sql)
  users := Users{}

  for row.Next() {
    user := User{}
    row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
    users = append(users, user)
  }
  return users
}

func LoginUser(username, password string) bool {
  user := GetUser("username", username)
  err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password));
  return err == nil
}




