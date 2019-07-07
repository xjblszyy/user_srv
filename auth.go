package user

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"goa.design/goa/v3/security"
	"time"
	. "user/utils"
	"user/utils/db/model"
)

// Credentials are invalid
type Unauthorized string

// Token scopes are invalid
type InvalidScopes string

// Email is Exist
type EmailExist string

// Email or Password Error
type EmailOrPasswordError string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e InvalidScopes) Error() string {
	return "Token scopes are invalid"
}

// ErrorName returns "invalid-scopes".
func (e InvalidScopes) ErrorName() string {
	return "invalid-scopes"
}

// Error returns an error description.
func (e EmailOrPasswordError) Error() string {
	return "Email or Password Error"
}

// Error returns an error description.
func (e EmailExist) Error() string {
	return "Email is exist"
}

var (
	// ErrUnauthorized is the error returned by Login when the request credentials
	// are invalid.
	ErrUnauthorized error = Unauthorized("invalid username and password combination")

	// ErrInvalidToken is the error returned when the JWT token is invalid.
	ErrInvalidToken error = Unauthorized("invalid token")

	// ErrInvalidTokenScopes is the error returned when the scopes provided in
	// the JWT token claims are invalid.
	ErrInvalidTokenScopes error = InvalidScopes("invalid scopes in token")

	ErrEmailExist error = EmailExist("email is exist")

	ErrEmailOrPassword error = EmailOrPasswordError("email or password error")

	// Key is the key used in JWT authentication
	Key = []byte(YamlConfig.Jwt.Secret)
)

// JWTAuth implements the authorization logic for service "userProfile" for the
// "jwt" security scheme.
func (s *userProfilesrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) { return Key, nil })
	if err != nil {
		return ctx, ErrInvalidToken
	}
	return ctx, nil
}

func GenToken(Id uint, Email string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": Id,
		"email": Email,
		"exp": time.Now().Add(time.Hour*24*7).Unix(),
	})
	return token.SignedString(Key)

}

func TokenToUser(Token string) (*model.User, error){
	type MyCustomClaims struct {
		Id int
		Email string
		jwt.StandardClaims
	}
	var user *model.User
	token, err := jwt.ParseWithClaims(Token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		user, err = db.FindUserById(claims.Id)
		if err != nil{
			return user, err
		}else{
			return user, nil
		}
	} else {
		return user, err
	}
}

