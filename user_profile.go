package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	userprofile "user/gen/user_profile"
	"user/utils"
	"user/utils/db/pg"
	"user/utils/db/redis"
	"user/utils/email"
)

var db pg.UserService = pg.UserService{}

// userProfile service example implementation.
// The example methods log the requests and return zero values.
type userProfilesrvc struct {
	logger *log.Logger
}

// NewUserProfile returns the userProfile service implementation.
func NewUserProfile(logger *log.Logger) userprofile.Service {
	return &userProfilesrvc{logger}
}

// Show userProfile by Token
func (s *userProfilesrvc) Retrieve(ctx context.Context, p *userprofile.RetrievePayload) (res *userprofile.UserProfile, err error) {
	res = &userprofile.UserProfile{}
	s.logger.Print("userProfile.retrieve")
	User, err := TokenToUser(p.Token)
	if err != nil {
		return &userprofile.UserProfile{}, err
	}
	user, err := pgCnn.FindUserById(int(User.ID))
	if err != nil {
		return &userprofile.UserProfile{}, err
	} else if user.IsActivate == false {
		err = errors.New("you email is not active, pleace activate email first")
		return &userprofile.UserProfile{}, err
	}else {
		return &userprofile.UserProfile{
			ID:       strconv.Itoa(int(user.ID)),
			Email:    user.Email,
			Avator:   &user.Avatar,
			Nickname: &user.Nickname,
		}, nil
	}
}

// Add new user and return its ID.
func (s *userProfilesrvc) Create(ctx context.Context, p *userprofile.AddUser) (res string, err error) {
	user, err := pgCnn.Register(p.Email, p.Password)
	if err != nil {
		return "", ErrEmailExist
	} else {
		Code := fmt.Sprintf("%v%v", utils.GetRandomString(16), user.ID)
		fmt.Println(Code)
		Expira := time.Duration(10 * time.Minute)
		err = redis.Set(Code, strconv.Itoa(int(user.ID)), Expira)
		if err != nil {
			s.logger.Println("redis set key and value error: ", err)
		}
		Url := fmt.Sprintf("http://0.0.0.0:8000/email/active/%s", Code)
		go email.SendToEmail(p.Email, Url)
		return "register success, please return you email address and active you email", nil
	}
}

// Creates a valid JWT
func (s *userProfilesrvc) Signin(ctx context.Context, p *userprofile.SigninPayload) (res string, err error) {
	s.logger.Print("userProfile.signin")
	user, err := pgCnn.FindOneUser(p.Email, p.Password)
	if err != nil || user.ID == 0 {
		return "", ErrEmailOrPassword
	} else if user.IsActivate == false{
		err = errors.New("you email is not active, pleace activate email first")
		return "", err
	} else {
		token, err := GenToken(user.ID, p.Email)
		if err != nil {
			log.Println("create token error: ", err)
		}
		return token, nil
	}
}

// Update avatar and nickname about user
func (s *userProfilesrvc) Update(ctx context.Context, p *userprofile.UpdateUser) (err error) {
	s.logger.Print("userProfile.update")
	user, err := TokenToUser(p.Token)
	if err != nil {
		return err
	}else if user.IsActivate == false{
		err = errors.New("you email is not active, pleace activate email first")
		return err
	} else{
		err = pgCnn.UserUpdateById(p.Avator, p.Nickname, int(user.ID))
		if err != nil{
			return err
		}else{
			return nil
		}
	}
}
