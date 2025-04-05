package model

type User struct {
	ID   int
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetID(id int) {
	u.ID = id
}
