package b

type User struct {
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

func GetUser() (*User, error) {
	return nil, nil
}

func NilObjFunc() {
	user, err := GetUser()
	if nil != err {
		print("balabala")
	}
	a := true
	if a {
		u := user.GetName() //may panic
		print(u)
	}

}
