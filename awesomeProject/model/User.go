package model

type User struct {
	UserName string
	PassWord string
	Email    string
}

// NewUser creates a new User instance and returns a pointer to it.
func NewUser(userName, password, email string) *User {
	u := &User{
		UserName: userName,
		PassWord: password,
		Email:    email,
	}
	return u
}

//// GetID returns the UserName of the User.
//func (u *User) GetID() int {
//	return u.ID
//}
//
//// SetID sets the UserName of the User to a new value.
//func (u *User) SetID(newValue string) {
//	u.UserName = newValue
//}
