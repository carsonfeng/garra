package a

type User struct {
	Name string
}

type Service struct {
}

func (svc *Service) User() *User {
	return nil
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetUser(uid int) (*User, int, error) {
	return nil, 0, nil
}

func (u *User) GetUserCache(uid int) (*User, int, error) {
	return nil, 0, nil
}

var svc = Service{}

func ForGetUser() {
	for i := 0; i < 100; i++ {
		svc.User().GetUser(i)
	}
}

func ForGetUserCache() {
	test := []int{1, 2, 3}
	for _, item := range test {
		i := item
		svc.User().GetUserCache(i)
	}
}
