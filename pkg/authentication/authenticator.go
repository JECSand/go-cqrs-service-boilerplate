package authentication

import (
	"context"
	"errors"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/enums"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Config settings for auth
type Config struct {
	SessionDuration     int `mapstructure:"sessionDuration"`     // 1
	IntegrationDuration int `mapstructure:"integrationDuration"` // 4380
}

func NewAuthConfig(uDur int, iDur int) *Config {
	return &Config{SessionDuration: uDur, IntegrationDuration: iDur}
}

type Authenticator interface {
	Authorize(ctx context.Context, method string) error
}

// authenticator
type authenticator struct {
	log         logging.Logger
	blacklist   BlacklistService
	accessRules map[string]enums.Role
}

// NewAuthenticator constructs a new authenticator
func NewAuthenticator(log logging.Logger, aRules map[string]enums.Role) *authenticator {
	return &authenticator{
		log:         log,
		accessRules: aRules,
	}
}

// Authorize a gRPC request
func (i *authenticator) Authorize(ctx context.Context, method string) error {
	accessRule, ok := i.accessRules[method]
	if !ok {
		return nil // unprotected endpoint
	}
	accessToken, err := GetTokenFromContext(ctx)
	if err != nil {
		return err
	}
	var authorized bool
	session, err := i.getTokenSession(accessToken)
	if err != nil {
		return err
	}
	switch accessRule {
	case enums.ROOT:
		if session.RootAdmin {
			authorized = true
		}
	case enums.MEMBER:
		authorized = true
	}
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}
	if authorized {
		return nil
	}
	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}

// getTokenSession validates & decrypts a JWT token, then returns the Session
func (i *authenticator) getTokenSession(accessToken string) (*Session, error) {
	if i.blacklist.CheckTokenBlacklist(accessToken) {
		return nil, errors.New("invalid token")
	}
	tokenSession, err := DecryptToken(accessToken)
	if err != nil {
		return nil, err
	}
	return tokenSession, nil
}
