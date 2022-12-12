package enums

// SessionType enumerates the potential values for Token.Type
type SessionType int

const (
	USER SessionType = iota
	INTEGRATION
)

// RoleClass enumerates the potential values for User.Role
type RoleClass int

const (
	MEMBER RoleClass = iota
	ADMIN
	ROOT
)
