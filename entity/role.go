package entity

type Role uint8

const (
	UserRole = 1 + iota
	AdminRole
)

func (r Role) String() string {
	switch r {
	case AdminRole:
		return "admin"

	case UserRole:
		return "user"
	}
	return ""
}
