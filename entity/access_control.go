package entity

// AccessControl only keeps allowed permissions
type AccessControl struct {
	ID           uint
	ActorID      uint
	ActorType    ActorType
	PermissionID uint
}

type ActorType string

const (
	RoleActorType = "role" //set access for all users with RoleID
	UserActorType = "user" //set access for user with UserID
)
