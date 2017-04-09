package models

const UserTable string = "users"

type User struct{
	Id			  int64		`json:"id"`
	Username 	string	`json:"username"`
	Password 	string	`json:"password"`
  Email     string  `json:"email"`
  errors    error
}

type Users []User

const UserSchema string = `CREATE TABLE users (
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(30) NOT NULL UNIQUE,
        password VARCHAR(30) NOT NULL,
        email VARCHAR(50),
        created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

func (this *User) Save(){
  if this.Id > 0{
    this.update()
  }else{
    this.insert()
  }
}

/*No Golang Way*/
func (this *User) Delete(){
  sql := "DELETE FROM users WHERE id=?"
  _, this.errors = modifyData(sql, this.Id)
}

func (this *User) update(){
  sql := "UPDATE users SET username=?,password=?,email=? WHERE id=?"
  _, this.errors = modifyData(sql, this.Username, this.Password, this.Email, this.Id)
}

func (this *User) insert(){
  sql := "INSERT users SET username=?,password=?,email=?"
  this.Id, this.errors = insertData(sql, this.Username, this.Password, this.Email)
}

func (this *User) SetPassword(password string){
  this.Password = password
}

func (this *User) GetErrors() error{
  return this.errors
}

func NewUser(username, password, email string) User{
  user := User{Username: username, Email: email, errors: nil}
  user.SetPassword(password)
  return user 
}

func CreateUser(username, password, email string) User{
  user := NewUser(username, password, email)
  user.Save()
  return user
}

func GetUser(id int) User{
  sql := "SELECT id, username, email FROM users WHERE id=?"
  user := User{}
  row, _ := query(sql, id)
  for row.Next() {
    row.Scan(&user.Id, &user.Username, &user.Email)
  }
  return user
}

func GetUsers() Users{
  sql := "SELECT id, username, email FROM users"
  row, _ := query(sql)
  users := Users{}

  for row.Next() {
    user := User{}
    row.Scan(&user.Id, &user.Username, &user.Email)
    users = append(users, user)
  }
  return users
}



