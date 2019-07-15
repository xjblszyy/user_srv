package app

import (
	"context"
	"fmt"
	"github.com/json-iterator/go"
	"log"
	"strconv"
	"user-srv/gen/user"
	pg "user-srv/utils/db"
	"user-srv/utils/email"
	"user-srv/utils/redis"
)


var db pg.UserService = pg.UserService{}
var json = jsoniter.ConfigCompatibleWithStandardLibrary


// user service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
}

type userRetrieve struct{
	Id uint `json:"id"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Nickname string `json:"nickname"`
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{logger}
}

// Show user info by Token
func (s *usersrvc) Retrieve(ctx context.Context, p *user.RetrievePayload) (res *user.ResponseData, err error) {
	s.logger.Print("userProfile.retrieve")
	userIns, err := TokenToUser(p.Token)
	if err != nil {
		return res, user.MakeInvalideToken(fmt.Errorf("token invalide"))
	}
	if userIns.IsActivate == false {
		return res, user.MakeNotFound(fmt.Errorf("you email is not active, pleace activate email first"))
	}

	u := userRetrieve{
		Id:userIns.ID,
		Email:userIns.Email,
		Avatar:userIns.Avatar,
		Nickname:userIns.Nickname,
	}
	jsonStr, _ := json.Marshal(u)
	res = &user.ResponseData{Code:200,
		Message:"success",
		Data:string(jsonStr)}
	return res, nil
}

// Add new user
func (s *usersrvc) Create(ctx context.Context, p *user.AddUser) (res *user.ResponseData, err error) {
	userIns, err := db.Register(p.Email, p.Password)
	if err != nil {
		return res, user.MakeEmailExist(fmt.Errorf("email is exist"))
	}

	email.PreSendEmail(int(userIns.ID), userIns.Email)
	res = &user.ResponseData{Code:200,
		Message:"success",
		Data:"regist success, please return your email address and active your email"}
	return res, nil
}

// Creates a valid JWT
func (s *usersrvc) Signin(ctx context.Context, p *user.Signin) (res *user.ResponseData, err error) {
	s.logger.Print("userProfile.signin")
	userIns, err := db.FindOneUser(p.Email, p.Password)
	if err != nil || userIns.ID == 0 {
		return &user.ResponseData{}, user.MakeNotFound(fmt.Errorf("user not found"))
	}
	if userIns.IsActivate == false{
		err = user.MakeCodeInvalide(fmt.Errorf("you email is not active, pleace activate email first"))
		return &user.ResponseData{}, err
	}

	token, err := GenToken(userIns.ID, p.Email)
	if err != nil {
		log.Println("create token error: ", err)
	}

	data:= map[string]string{"token":token}
	jsonStr, _ := json.Marshal(data)
	reData := &user.ResponseData{
		Code:200,
		Message:"success",
		Data:string(jsonStr),
	}
	return reData, nil

}

// Update avatar and nickname about user
func (s *usersrvc) Update(ctx context.Context, p *user.UpdateUser) (err error) {
	s.logger.Print("userProfile.update")
	userIns, err := TokenToUser(p.Token)
	if err != nil {
		return err
	}
	if userIns.IsActivate == false{
		err = user.MakeCodeInvalide(fmt.Errorf("you email is not active, pleace activate email first"))
		return err
	}

	err = db.UserUpdateById(p.Avatar, p.Nickname, int(userIns.ID))
	if err != nil{
		return err
	}

	return nil
}

// Send email to active user
func (s *usersrvc) SendEmail(ctx context.Context, p *user.SendEmail) (res *user.ResponseData, err error) {
	userIns, err := db.FindUserByEmail(p.Email)
	if err != nil{
		return res, user.MakeEmailNotFound(fmt.Errorf("email not found or email is avtivated"))
	}

	if userIns.ID > 0{
		email.PreSendEmail(int(userIns.ID), userIns.Email)
		reData := &user.ResponseData{
			Code:200,
			Message:"success",
			Data:"send email success, please return your email address and active your email",
		}
		return reData, nil
	}

	return res, user.MakeEmailNotFound(fmt.Errorf("email not found"))
}

// Active email to user
func (s *usersrvc) ActiveEmail(ctx context.Context, p *user.EmailCode) (res *user.ResponseData, err error) {
	res = &user.ResponseData{}
	id, err := redis.Get(p.Code)
	if err != nil{
		return res, user.MakeCodeInvalide(fmt.Errorf("code not found or is expired"))
	}

	_ = redis.Delete(p.Code)
	Id, err := strconv.Atoi(id)
	if err != nil{
		log.Println("string to int error: ", err)
	}

	_, err = db.ActivateUserById(Id)
	if err != nil{
		return res, user.MakeNotFound(fmt.Errorf("user not found"))
	}

	res.Data = "active success"
	res.Code = 200
	res.Message = "success"
	return res, nil
}
