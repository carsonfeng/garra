package a

type User struct {
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

func GetUser() (*User, int, error) {
	return nil, 0, nil
}

func NilObjFunc2() {
	var temp string
	user2, t, err := GetUser()
	if nil != err {
		//return
		print("error" + err.Error())
	}
	_ = t
	temp = user2.Name
	temp = user2.GetName()
	user3, t2, err2 := GetUser()
	if nil != err2 {
		print("error" + err.Error())
	}
	user3.GetName()
	_ = t2
	temp = user3.GetName()
	temp = user3.Name

	_ = temp
	_ = user2.GetName()
	_ = user3.GetName()
}

//func NilObjFunc2() {
//	for i := 0; i < 3; i++ {
//		user4, t4, err4 := GetUser()
//		if nil != err4 {
//			//fmt.Printf("err: %v", err)
//			t := 1 + 1
//			_ = t
//			continue
//		}
//		name := user4.GetName()
//		_ = name
//		_ = user4
//		_ = t4
//	}
//}

func cc() {

}
