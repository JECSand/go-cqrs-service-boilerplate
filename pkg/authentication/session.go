package authentication

import (
	"errors"
	"fmt"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/enums"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/utilities"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// Session stores the structured data from a session token for use
type Session struct {
	UserId     string
	RootAdmin  bool
	Type       enums.SessionType
	Expiration int64
	secret     string
}

// NewSession constructions a new Token
func NewSession(userId string, root bool, tokenType enums.SessionType, secretStr string) *Session {
	return &Session{
		UserId:     userId,
		RootAdmin:  root,
		Type:       tokenType,
		Expiration: 0,
		secret:     secretStr,
	}
}

// setExpiration returns the unix time for token expiration
func (t *Session) setExpiration() {
	var duration int
	switch {
	case t.Type == enums.USER:
		duration = utilities.StringToInt(os.Getenv("USER_DURATION"))
	case t.Type == enums.INTEGRATION:
		duration = utilities.StringToInt(os.Getenv("INTEGRATION_DURATION"))
	default:
		duration = 1
	}
	t.Expiration = time.Now().Add(time.Hour * time.Duration(duration)).Unix()
}

// NewToken is used to create an encrypted token string from an auth Session
func (t *Session) NewToken() (string, error) {
	if t.UserId == "" {
		return "", errors.New("missing required token claims")
	}
	t.setExpiration()
	if t.Expiration == 0 {
		return "", errors.New("new token must have a expiration time greater than 0")
	}
	var MySigningKey = []byte(t.secret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = t.UserId
	claims["root"] = t.RootAdmin
	claims["type"] = t.Type
	claims["exp"] = t.Expiration
	return token.SignedString(MySigningKey)
}

// DecryptToken a Session from an encrypted token string
func DecryptToken(tokenStr string) (*Session, error) {
	var session Session
	if tokenStr == "" {
		return &session, errors.New("unauthorized")
	}
	var MySigningKey = []byte(os.Getenv("TOKEN_SECRET"))
	parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return MySigningKey, nil
	})
	if err != nil {
		return &session, err
	}
	if parsedToken.Valid {
		tokenClaims := parsedToken.Claims.(jwt.MapClaims)
		session.UserId = tokenClaims["id"].(string)
		session.RootAdmin = tokenClaims["root"].(bool)
		session.RootAdmin = tokenClaims["type"].(bool)
		session.Expiration = tokenClaims["exp"].(int64)
		return &session, nil
	}
	return &session, errors.New("invalid token")
}
