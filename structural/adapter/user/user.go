package user

type User struct {
	ID       string
	Username string
	Email    string
}

type Store interface {
	Insert(User) (string, error)
}

func CreateUser(s Store, u User) (ID string, err error) {
	ID, err = s.Insert(u)
	if err != nil {
		return "", err
	}
	return ID, nil
}
