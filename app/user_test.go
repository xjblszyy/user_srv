package app

import (
	"bytes"
	"context"
	"log"
	"reflect"
	"testing"
	"time"
	"user-srv/gen/user"
	"user-srv/utils/redis"
)

var (
	logBuf bytes.Buffer
	token  string
	email1 string = "1366593680001@qq.com"
	email2 string = "1366593680002@qq.com"
	email3 string = "1366593680003@qq.com"
	password string = "123456"
)
var newLog *log.Logger = log.New(&logBuf, "", log.Ltime)

//func TestNewUser(t *testing.T) {
//	var newLog *log.Logger = log.New(&logBuf, "", log.Ltime)
//	type args struct {
//		logger *log.Logger
//	}
//	tests := []struct {
//		name string
//		args args
//		want user.Service
//	}{
//		{name:"newUser1", args:args{newLog}, want:nil},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewUser(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewUser() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func test_usersrvc_Create(t *testing.T) {

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		ctx context.Context
		p   *user.AddUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *user.ResponseData
		wantErr bool
	}{
		{"create1", fields{newLog}, args{ctx: context.Background(), p: &user.AddUser{Email: email1, Password: password}}, &user.ResponseData{Code: 200, Message: "success", Data: "regist success, please return your email address and active your email"}, false},
		{"create2", fields{newLog}, args{ctx: context.Background(), p: &user.AddUser{Email: email1, Password: password}}, nil, true},
		{"create3", fields{newLog}, args{ctx: context.Background(), p: &user.AddUser{Email: email2, Password: password}}, &user.ResponseData{Code: 200, Message: "success", Data: "regist success, please return your email address and active your email"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &usersrvc{
				logger: tt.fields.logger,
			}
			gotRes, err := s.Create(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersrvc.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("usersrvc.Create() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func test_usersrvc_ActiveEmail(t *testing.T) {

	var code string = redis.GetEmailCode()
	type fields struct {
		logger *log.Logger
	}
	type args struct {
		ctx context.Context
		p   *user.EmailCode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *user.ResponseData
		wantErr bool
	}{
		{"active1", fields{newLog}, args{ctx: context.Background(), p: &user.EmailCode{Code: "123456"}}, &user.ResponseData{}, true},
		{"active2", fields{newLog}, args{ctx: context.Background(), p: &user.EmailCode{Code: code}}, &user.ResponseData{Code: 200, Message: "success", Data: "active success"}, false},
	}
	for _, tt := range tests {
		time.Sleep(time.Second * 1)
		t.Run(tt.name, func(t *testing.T) {
			s := &usersrvc{
				logger: tt.fields.logger,
			}
			gotRes, err := s.ActiveEmail(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersrvc.ActiveEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("usersrvc.ActiveEmail() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func test_usersrvc_SendEmail(t *testing.T) {
	type fields struct {
		logger *log.Logger
	}
	type args struct {
		ctx context.Context
		p   *user.SendEmail
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *user.ResponseData
		wantErr bool
	}{
		{"sendEmail2", fields{newLog}, args{ctx: context.Background(), p: &user.SendEmail{Email: email2}}, &user.ResponseData{Code: 200, Message: "success", Data: "send email success, please return your email address and active your email"}, false},
		{"sendEmail1", fields{newLog}, args{ctx: context.Background(), p: &user.SendEmail{Email: email1}}, nil, true},
		{"sendEmail3", fields{newLog}, args{ctx: context.Background(), p: &user.SendEmail{Email: email3}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &usersrvc{
				logger: tt.fields.logger,
			}
			gotRes, err := s.SendEmail(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersrvc.SendEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("usersrvc.SendEmail() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func test_usersrvc_Signin(t *testing.T) {
	var newLog *log.Logger = log.New(&logBuf, "", log.Ltime)
	type fields struct {
		logger *log.Logger
	}
	type args struct {
		ctx context.Context
		p   *user.Signin
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *user.ResponseData
		wantErr bool
	}{
		{"signin2", fields{newLog}, args{ctx: context.Background(), p: &user.Signin{Email: email2, Password: password}}, &user.ResponseData{}, true},
		{"signin3", fields{newLog}, args{ctx: context.Background(), p: &user.Signin{Email: email3, Password: password}}, &user.ResponseData{}, true},
		{"signin1", fields{newLog}, args{ctx: context.Background(), p: &user.Signin{Email: email1, Password: password}}, &user.ResponseData{Code: 200, Message: "success", Data: "{\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjEzNjY1OTM2ODAwMDFAcXEuY29tIiwiZXhwIjoxNTYzNTE5OTk3LCJpZCI6MjE4fQ.tiwhe3Lp23zWJbZ8yCnPtEIcqMuPn02VixLduZXKZnY\"}"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &usersrvc{
				logger: tt.fields.logger,
			}
			gotRes, err := s.Signin(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersrvc.Signin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.wantRes = gotRes
			var mapResult map[string]string
			_ = json.Unmarshal([]byte(gotRes.Data), &mapResult)
			token = mapResult["token"]
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("usersrvc.Signin() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func test_usersrvc_Retrieve(t *testing.T) {
	var newLog *log.Logger = log.New(&logBuf, "", log.Ltime)
	t.Log(token)
	type fields struct {
		logger *log.Logger
	}
	type args struct {
		ctx context.Context
		p   *user.RetrievePayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *user.ResponseData
		wantErr bool
	}{
		{"retrieve1", fields{newLog}, args{ctx: context.Background(), p: &user.RetrievePayload{token}}, &user.ResponseData{Code: 200, Message: "success", Data: "regist success, please return your email address and active your email"}, false},
		{"retrieve2", fields{newLog}, args{ctx: context.Background(), p: &user.RetrievePayload{"123"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &usersrvc{
				logger: tt.fields.logger,
			}
			gotRes, err := s.Retrieve(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersrvc.Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("usersrvc.Retrieve() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func test_usersrvc_Update(t *testing.T) {
	type fields struct {
		logger *log.Logger
	}
	t.Log(token)
	type args struct {
		ctx context.Context
		p   *user.UpdateUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"update1", fields{newLog}, args{ctx:context.Background(),p:&user.UpdateUser{Avatar:"123", Nickname:"123", Token:token}}, false},
		{"update2", fields{newLog}, args{ctx:context.Background(),p:&user.UpdateUser{Avatar:"123", Nickname:"123", Token:"123456789"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &usersrvc{
				logger: tt.fields.logger,
			}
			if err := s.Update(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("usersrvc.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAll(t *testing.T) {
	t.Run("test_usersrvc_Create", test_usersrvc_Create)
	t.Run("test_usersrvc_ActiveEmail", test_usersrvc_ActiveEmail)
	t.Run("test_usersrvc_SendEmail", test_usersrvc_SendEmail)
	t.Run("test_usersrvc_Signin", test_usersrvc_Signin)
	t.Run("test_usersrvc_Retrieve", test_usersrvc_Retrieve)
	//t.Run("test_usersrvc_Update", test_usersrvc_Update)
}


