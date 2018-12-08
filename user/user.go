package user

type User struct {
	counter  int
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

func NewUser() *User {
	return &User{
		counter: 1,
	}
}

func (uc *UserCollection) AddUser(user *User) {
	uc.Users[uc.counter] = user
	uc.counter = uc.counter + 1
}

type UserCollection struct {
	counter int
	Users   map[int]*User
}

func NewUserCollection() *UserCollection {
	return &UserCollection{
		Users: map[int]*User{},
	}
}
