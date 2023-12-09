package postgres

import (
	"database/sql"
	"github.com/MehmetTalhaSeker/design-patterns/structural/adapter/user"
)

type Store struct {
	db *sql.DB
}

func (p *Store) Insert(u user.User) (ID string, err error) {
	// save to db logic
	return u.ID, nil
}
