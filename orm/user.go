package models

const UserTable string = "users"

type User struct{
  Id        int64   `json:"id"`
  Username  string  `json:"username"`
  Password  string  `json:"password"`
  Email     string  `json:"email"`
}

type Users []User

func (this *User) Save() bool{
  if this.Id > 0{
    return this.update()
  }else{
    return this.insert()
  }
}

func (this *User) Delete(){
  db.Delete(&this)
}

func (this *User) update() bool {
  db.Save(&this)
  return true
}

func (this *User) insert() bool {
  db.Create(&this)
  return true
}

func (this *User) SetPassword(password string){
  this.Password = password
}

func NewUser(username, password, email string) *User{
  user :=&User{Username: username, Email: email }
  user.SetPassword(password)
  return user 
}

func CreateUser(username, password, email string) *User{
  user := NewUser(username, password, email)
  user.Save()
  return user
}
  
func GetUser(id int) User{
  user := User{}
  db.Where("id = ?", id).First(&user)
  return user
}

func GetUsers() Users{
  users := Users{}
  db.Find(&users)
  return users
}



