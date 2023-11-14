package data

const UserID = "userID"

const (
	RoleUser = iota + 1
	RoleAdmin
)

//go:generate reform

// User represents a user in our system.
//
//go:generate easyjson -all data.go
//reform:users
type User struct {
	ID int64 `json:"-" reform:"id,pk"`

	Login    string `json:"Login" reform:"login"`
	Password string `json:"Password" reform:"password"`

	Name string `json:"Name" reform:"name"`
	Role int    `json:"-" reform:"role_id"`

	Deleted bool `json:"-" reform:"deleted"`
}
