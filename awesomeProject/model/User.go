package model

type User struct {
	ID       int
	UserName string
	FullName string
	Email    string
}

// NewUser creates a new User instance and returns a pointer to it.
func NewUser(userName, fullName, email string, Id int) *User {
	u := &User{
		ID:       Id,
		UserName: userName,
		FullName: fullName,
		Email:    email,
	}
	return u
}

// GetID returns the UserName of the User.
func (u *User) GetID() int {
	return u.ID
}

// SetID sets the UserName of the User to a new value.
func (u *User) SetID(newValue string) {
	u.UserName = newValue
}
