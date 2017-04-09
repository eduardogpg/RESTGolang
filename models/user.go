package models

const UserTable string = "users"

type User struct{
	Id			  int64		 `json:"id"`
	Username 	string	`json:"username"`
	Password 	string	`json:"password"`
  Email     string  `json:"email"`
}

type Users []User

const UserSchema string = `CREATE TABLE users (
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(30) NOT NULL,
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
  
func (this *User) Delete(){
  sql := "DELETE FROM users WHERE id=?"
  executeSql(sql, this.Id)
}

func (this *User) update(){
  sql := "UPDATE users SET username=?,password=?,email=? WHERE id=?"
  executeSql(sql, this.Username, this.Password, this.Email, this.Id)
}

func (this *User) insert(){
  sql := "INSERT users SET username=?,password=?,email=?"
  this.Id = insertData(sql, this.Username, this.Password, this.Email)
}

func NewUser(username, password, email string) User{
  user := User{Username: username, Email: email}
  user.Password = password
  return user 
}

func CreateUser(username, password, email string) User{
  user := NewUser(username, password, email)
  return user
}

func GetUser(id int) User{
  sql := "SELECT id, username, email FROM users WHERE id=?"
  user := User{}
  row := executeQuery(sql, id)
  for row.Next() {
    row.Scan(&user.Id, &user.Username, &user.Email)
  }
  return user
}

func GetUsers() Users{
  sql := "SELECT id, username, email FROM users"
  row := executeQuery(sql)
  users := Users{}

  for row.Next() {
    user := User{}
    row.Scan(&user.Id, &user.Username, &user.Email)
    users = append(users, user)
  }
  return users
}



