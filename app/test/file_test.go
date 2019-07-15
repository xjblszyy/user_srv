package test

//import (
//	"context"
//	"log"
//	"reflect"
//	"testing"
//	file "user-srv/gen/file"
//)
//
//func TestNewFile(t *testing.T) {
//	type args struct {
//		logger *log.Logger
//	}
//	tests := []struct {
//		name string
//		args args
//		want file.Service
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewFile(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewFile() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_filesrvc_Upload(t *testing.T) {
//	type fields struct {
//		logger *log.Logger
//	}
//	type args struct {
//		ctx context.Context
//		p   *file.UploadPayload
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantRes *file.ResponseData
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &filesrvc{
//				logger: tt.fields.logger,
//			}
//			gotRes, err := s.Upload(tt.args.ctx, tt.args.p)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("filesrvc.Upload() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotRes, tt.wantRes) {
//				t.Errorf("filesrvc.Upload() = %v, want %v", gotRes, tt.wantRes)
//			}
//		})
//	}
//}