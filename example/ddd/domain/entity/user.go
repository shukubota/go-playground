package entity

type User struct {
	ID   *uint
	Name string
}

func NewUser(i *uint, n string) *User {
	return &User{
		ID:   i,
		Name: n,
	}
}

func (u *User) ChangeName(n string) error {
	u.Name = n
	return nil
}
