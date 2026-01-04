package domain

type User struct {
	id       int
	username string
	password string
	admin    bool
}

type NewUserData struct {
	ID       int
	Username string
	Password string
	Admin    bool
}

func NewUser(data NewUserData) (User, error) {
	return User{
		id:       data.ID,
		username: data.Username,
		password: data.Password,
		admin:    data.Admin,
	}, nil
}

func (u User) ID() int {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func (u User) Password() string {
	return u.password
}

func (u User) Admin() bool {
	return u.admin
}
