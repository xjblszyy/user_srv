package app

//import (
//	"context"
//	"fmt"
//
//	"goa.design/goa/v3/security"
//)
//
//// JWTAuth implements the authorization logic for service "user" for the "jwt"
//// security scheme.
//func (s *usersrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
//	//
//	// TBD: add authorization logic.
//	//
//	// In case of authorization failure this function should return
//	// one of the generated error structs, e.g.:
//	//
//	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
//	//
//	// Alternatively this function may return an instance of
//	// goa.ServiceError with a Name field value that matches one of
//	// the design error names, e.g:
//	//
//	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
//	//
//	return ctx, fmt.Errorf("not implemented")
//}


import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"goa.design/goa/v3/security"
	"time"
	. "user-srv/utils"
	. "user-srv/utils/db"
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
func (s *usersrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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
		"exp": time.Now().Add(time.Hour*7*24).Unix(),
	})
	return token.SignedString(Key)

}

func TokenToUser(Token string) (*User, error){
	type MyCustomClaims struct {
		Id int
		Email string
		jwt.StandardClaims
	}
	var user *User
	token, err := jwt.ParseWithClaims(Token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		user, err = db.FindUserById(claims.Id)
		if err != nil{
			return user, err
		}

		return user, nil
	}

	return user, err
}