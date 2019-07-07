package user

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"user/gen/email"
	"user/utils"
	"user/utils/db/pg"
	"user/utils/db/redis"
	emailClient "user/utils/email"
)

var pgCnn pg.UserService = pg.UserService{}

// email service example implementation.
// The example methods log the requests and return zero values.
type emailsrvc struct {
	logger *log.Logger
}

// NewEmail returns the email service implementation.
func NewEmail(logger *log.Logger) email.Service {
	return &emailsrvc{logger}
}

// Active user by email code
func (s *emailsrvc) Active(ctx context.Context, p *email.ActivePayload) (res string, err error) {
	s.logger.Print("email.active")
	id, err := redis.Get(p.Code)
	if err != nil{
		return "", err
	}else{
		Id, err := strconv.Atoi(id)
		if err != nil{
			log.Println("string to int error: ", err)
		}
		err = pgCnn.ActivateUserById(Id)
		if err != nil{
			return "", err
		}else{
			return "激活成功", nil
		}
	}
}

// Send email to active user
func (s *emailsrvc) SendEmail(ctx context.Context, p *email.SendEmailPayload) (res string, err error) {
	s.logger.Print("email.send_email")
	user, err := pgCnn.FindUserByEmail(p.Email)
	if err != nil{
		return "", err
	}else{
		if user.ID > 0{
			Code := fmt.Sprintf("%v%v", utils.GetRandomString(16), user.ID)
			fmt.Println(Code)
			Expira := time.Duration(10 * time.Minute)
			err = redis.Set(Code, strconv.Itoa(int(user.ID)), Expira)
			if err != nil {
				s.logger.Println("redis set key and value error: ", err)
			}
			Url := fmt.Sprintf("http://0.0.0.0:8000/email/active/%s", Code)
			go emailClient.SendToEmail(p.Email, Url)
			return "send active email success, please open your email", nil
		}else{
			return "user not found", nil
		}
	}
}
