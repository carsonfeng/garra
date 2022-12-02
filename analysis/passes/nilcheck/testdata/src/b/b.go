package b

type User struct {
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

func GetUserList() ([]User, error) {
	return nil, nil
}

func NilObjFunc() {
	userList, err := GetUserList()
	if nil != err {
		print("error" + err.Error())
	}
	len2 := len(userList) //it's correct！
	print(len2)
}
