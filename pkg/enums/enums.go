package enums

// SessionType enumerates the potential values for Token.Type
type SessionType int

const (
	USER SessionType = iota
	INTEGRATION
)

// Role enumerates the potential values for User.Role
type Role int

const (
	MEMBER Role = iota
	ADMIN
	ROOT
)
