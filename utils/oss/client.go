package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"log"
	"strconv"
	"time"
	. "user-srv/utils"
)

var Client *oss.Client

func InitOss(){
	client, err := oss.New(YamlConfig.Oss.Endpoint, YamlConfig.Oss.AccessSecretKey, YamlConfig.Oss.AccessSecretKey)
	if err != nil{
		log.Fatalln("conncet aliyun oss error: ", err)
	}else{
		Client = client
	}
}

func UploadFile(Name string, File io.Reader) (string, error) {
	bucket, err := Client.Bucket(YamlConfig.Oss.BucketName)
	if err != nil {
		return "", err
	}
	objectKey := strconv.Itoa(int(time.Now().Unix()))+Name
	err = bucket.PutObject(objectKey, File)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.%s/%s", YamlConfig.Oss.BucketName, YamlConfig.Oss.Endpoint, objectKey), nil
}