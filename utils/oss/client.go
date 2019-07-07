package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"log"
	"strconv"
	"time"
	. "user/utils"
)

var Client *oss.Client

var(
	endpoint = YamlConfig.Oss.Endpoint
	accessKeyId = YamlConfig.Oss.AccessKey
	accessKeySecret = YamlConfig.Oss.AccessSecretKey
	bucketName = YamlConfig.Oss.BucketName
)

func init(){
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil{
		log.Fatalln("conncet aliyun oss error: ", err)
	}else{
		Client = client
	}
}

func UploadFile(Name string, File io.Reader) (string, error) {
	bucket, err := Client.Bucket(bucketName)
	if err != nil {
		return "", err
	}else{
		objectKey := strconv.Itoa(int(time.Now().Unix()))+Name
		err = bucket.PutObject(objectKey, File)
		if err != nil {
			return "", err
		}else{
			return fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, objectKey), nil
		}
	}
}