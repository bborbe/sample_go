package user

import "fmt"

type user struct {
	Firstname string
	Lastname  string
}

func New() *user {
	return new(user)
}

func (a *user) DisplayName() string {
	return fmt.Sprintf("%s %s", a.Firstname, a.Lastname)
}
