package model

import (
	"errors"
	"time"
)

type User struct {
	Id        uint
	Name      string
	CreatedAt time.Time
}

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, errors.New("nameが空です")
	}

	user := &User{
		Name:      name,
		CreatedAt: time.Now(),
	}

	return user, nil
}
