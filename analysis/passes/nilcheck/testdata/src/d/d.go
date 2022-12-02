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
		panic("err")
	}
	u := user.GetName() //PASS. because panic in if.then
	print(u)
}
