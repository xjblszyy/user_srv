package email

import (
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"strings"
	. "user-srv/utils"
	"user-srv/utils/redis"
)

func sendToEmail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	if err != nil{
		return err
	}
	return nil
}

func SendToEmail(ToEmail, Url string){
	user := YamlConfig.Email.User
	password := YamlConfig.Email.Password
	host := YamlConfig.Email.Host
	to := ToEmail

	subject := YamlConfig.Email.Subject
	fmt.Println(Url) // TODO
	body := fmt.Sprintf(`<html><body><h3><a href="%s">点击链接激活邮箱</a></h3></body></html>`, Url)
	err := sendToEmail(user, password, host, to, subject, body, "html")
	if err != nil{
		log.Println("send email error: ", err)
	}
}


func PreSendEmail(Id int, Email string){
	Code := fmt.Sprintf("%v%v", GetRandomString(16), Id)
	Expira := YamlConfig.Email.EmailCodeExp
	err := redis.RedisCnn.Set(Code, strconv.Itoa(int(Id)), Expira)
	if err != nil {
		log.Println("redis set key and value error: ", err)
	}
	baseUrl := YamlConfig.Email.ActiveUrl
	Url := fmt.Sprintf(baseUrl + "/%s", Code)  // TODO
	go SendToEmail(Email, Url)
}