package test

import "testing"

//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"testing"
//)
//
//func testAddUser(t *testing.T) {
//
//	cases := []struct {
//		email    string
//		password string
//		code int
//		resp interface{}
//	}{
//		{"xjblszyy1@163.com", "123456", 200,map[string]interface{}{"code":200,"message":"success","data":""}},
//		{"xjblszyy1@163.com", "123456", 400, map[string]interface{}{}},
//	}
//	for _, tt := range cases{
//		client := &http.Client{}
//		req := fmt.Sprintf(`{"email":%s", "passwprd": %s}`, tt.email, tt.password)
//		req_new := bytes.NewBuffer([]byte(req))
//		request, _ := http.NewRequest("POST", "http://0.0.0.0:8000/user/register", req_new)
//		request.Header.Set("Content-type", "application/json")
//		response, _ := client.Do(request)
//		if response.StatusCode ==tt.code {
//			body, _ := ioutil.ReadAll(response.Body)
//			fmt.Println(string(body))
//		}
//	}
//}
//
//
//func TestAll(t *testing.T) {
//	t.Run("testAddUser", testAddUser)
//}

func TestMain(m *testing.M) {
	// 清空数据库
	fmt.Println("begin")
	m.Run()
	// 清空数据库
}
