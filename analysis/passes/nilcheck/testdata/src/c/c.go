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
	for i := 0; i < 10; i++ {
		user, err := GetUser()
		if nil != err {
			print("error" + err.Error())
		}
		u := user.GetName() //may panic
		print(u)
	}
}
