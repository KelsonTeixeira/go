package main

import "fmt"

// When
// 1. When we need to update state
// 2. Pointer is 8 bytes
// 3. if we not send pointer, we are sending a copy of the value, which can be expensive if the value is large
type User struct {
	email    string
	username string
	age      int
}

func (u User) Email() string {
	return u.email
}

func (u *User) updateEmail(email string) {
	u.email = email
}

func UpdateEmail(u *User, email string) {
	u.email = email
}

func main() {
	user := User{
		email: "email@email.com",
	}
	fmt.Println(user.Email())

	user.updateEmail("new@email.com") // this case, because of the way method was written, no need to pass the pointer
	fmt.Println(user.Email())

	UpdateEmail(&user, "newemail@email.com") // in this case, we need to pass the pointer. & sends the address

	fmt.Println(user.Email())
}
