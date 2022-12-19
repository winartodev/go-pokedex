package enum

type Role int64

const (
	Public Role = iota
	User
	Admin
)

// String() method adds the ability to return
// role constants as a string rather than an int
func (r Role) String() string {
	switch r {
	case User:
		return "user"
	case Admin:
		return "admin"
	}
	return ""
}
