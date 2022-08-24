package cmd

import (
	"time"
)

type User struct {
	Name       string
	Status     int
	LastAction time.Time
}

type AllStat struct {
	Users     []User
	UserCount int
	CurUser   User
	DoesExist bool
}
