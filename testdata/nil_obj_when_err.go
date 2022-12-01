package testdata

import "fmt"

type User struct {
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

func getUser() (*User, error) {
	return nil, fmt.Errorf("not found")
}

func NilObjFunc() {
	//f := 1 / 0
	//_ = f

	user, err := getUser()
	if nil != err {
		fmt.Printf("err: %v", err)
	}
	_ = user.GetName()
}
