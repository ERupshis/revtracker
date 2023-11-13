package data

import (
	"fmt"
)

const UserID = "userID"

const (
	RoleUser = iota + 1
	RoleAdmin
)

// ErrUserNotFound missing user in database.
var ErrUserNotFound = fmt.Errorf("user not found")

//go:generate reform

// User represents a user in our system.
//
//go:generate easyjson -all data.go
//reform:users
type User struct {
	Login    string `json:"login" reform:"login"`
	Password string `json:"password" reform:"password"`

	ID   int64  `json:"-" reform:"id,pk"`
	Name string `json:"-" reform:"name"`
	Role int    `json:"-" reform:"role_id"`
}
