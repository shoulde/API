package user

type User struct {
	ID       int    `json:"ID,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

func NewUser() *User {
	return &User{
		ID: 1,
	}
}

func (uc *UserCollection) AddUser(user *User) {
	uc.Users = append(uc.Users, user)
}

type UserCollection struct {
	counter int
	Users   []*User
}

func NewUserCollection() *UserCollection {
	return &UserCollection{
		Users: []*User{},
	}
}
