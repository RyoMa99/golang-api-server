package model

import (
	"time"
)

type User struct {
	Id        uint
	name      string
	CreatedAt time.Time
}
