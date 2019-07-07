# user_srv

## 使用方法
### 1:配置好golang1.12+环境
### 2:git clone https://github.com/xjblszyy/user_srv.git
### 3:安装三方库

```
cd user_srv
export GO111MODULE=on
export GOPROXY=https://goproxy.io
go get -u goa.design/goa/v3
go mod download
```

### 4:配置config.yaml中的参数

```
server:
  addr: "http://127.0.0.1:8000" 
  grpc: "grpc://localhost:8080"
  debug: true 

postgresql:
  host: "0.0.0.0"
  port: "5432"
  dbname: "iot"
  user: "iot"
  password: "123456"
  sslmode: "disable"
  maxOpenConns: 100
  maxIdleConns: 10
  debug : true

redis:
  addr: "0.0.0.0:6379"
  password: ""
  db : 0

oss:
  endpoint: "oss-cn-xxx.aliyuncs.com"
  accessKey: "aaa"
  accessSecretKey: "bbb"
  bucketName: "ccc-ddd"

email:
  user: "1@1.com"
  password: "eee"
  host: "smtp.fff.com:587"
  subject: "welcome use iot system, pleace active your email"

jwt:
  secret: "secret"
```
### 5:编译并运行文件

```
cd cmd/user
go build
./user
```
