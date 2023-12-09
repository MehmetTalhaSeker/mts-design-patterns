package testutils

import (
	"github.com/MehmetTalhaSeker/design-patterns/structural/adapter/user"
)

type TestStore struct {
}

func (p *TestStore) Insert(u user.User) (ID string, err error) {
	// easier tests
	return u.ID, nil
}
