package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	. "user-srv/utils"
)

type User struct {
	gorm.Model
	Email string `gorm:"size:63;not null;unique;index"`
	Password string `gorm:"size:63;not null"`
	Avatar string `gorm:"size:63"`
	Nickname string `gorm:"size:63;index"`
	IsActivate bool `gorm:"default:false;index"`
}

var PgCnn *gorm.DB

func InitDB(){
	var err error
	args := YamlConfig.Postgresql
	PgCnn, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		args.Host, args.Port, args.User, args.Dbname, args.Password, args.Sslmode))
	if err != nil {
		log.Fatalln("connect pg error: ", err)
	}
	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	PgCnn.DB().SetMaxIdleConns(YamlConfig.Postgresql.MaxIdleConns)
	PgCnn.DB().SetMaxOpenConns(YamlConfig.Postgresql.MaxOpenConns)

	// 启用Logger，显示详细日志
	PgCnn.LogMode(YamlConfig.Postgresql.Debug)
	PgCnn.AutoMigrate(&User{})
}




