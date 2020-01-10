package main

import (
	"fmt"

	p "github.com/girishg4t/golang-samples/stringer/painkiller"
)

type User struct {
	Name  string
	Email string
}

// String satisfies the fmt.Stringer interface for the User type
func (u User) String() string {
	return fmt.Sprintf("%s do it <%s>", u.Name, u.Email)
}

func main() {
	u := User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	fmt.Println(u.String())
	p1 := p.PillType{3}
	fmt.Println(p1.String())

}
