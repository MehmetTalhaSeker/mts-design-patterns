package user

type User struct {
	email    string
	username string
	password string
}

type Builder struct {
	User
}

func NewBuilder() *Builder {
	return &Builder{User: User{}}
}

func (u *Builder) SetEmail(email string) *Builder {
	u.email = email
	return u
}

func (u *Builder) SetUsername(username string) *Builder {
	u.username = username
	return u
}

func (u *Builder) SetPassword(password string) *Builder {
	u.password = password
	return u
}

func (u *Builder) Build() *User {
	return &u.User
}
