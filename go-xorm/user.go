package go_xorm

type User struct {
	ID      int64 `xorm:"pk autoincr"`
	Name    string
	Email   string
	Address Address `xorm:"jsonb"`
}

type Address struct {
	City    string
	Country string
}

func CreateUser(user *User) (*User, error) {
	_, err := en.Insert(user)
	return user, err
}

func ListUser() ([]User, error) {
	users := make([]User, 0)
	err := en.Find(&users, User{ID: 0})
	return users, err
}
