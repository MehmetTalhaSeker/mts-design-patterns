package user_test

import (
	"github.com/MehmetTalhaSeker/design-patterns/structural/adapter/testutils"
	"github.com/MehmetTalhaSeker/design-patterns/structural/adapter/user"
	"testing"
)

func TestCreate(t *testing.T) {
	s := &testutils.TestStore{}
	w := user.User{
		ID:       "test",
		Username: "test",
		Email:    "test",
	}
	want := "test"
	got, err := user.CreateUser(s, w)

	if err != nil {
		t.Errorf("smt went wrong: %v", err)
	}

	if want != got {
		t.Errorf("want: %s, got:%s", want, got)
	}
}
