package models

// define a user struct
type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

// get auser
func NewUser() *User {
	return &User{}
}

// get set methods
func (u *User) GetID() int {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetID(id int) *User {
	u.ID = id
	return u
}

func (u *User) SetName(name string) *User {
	u.Name = name
	return u
}

func (u *User) SetEmail(email string) *User {
	u.Email = email
	return u
}
