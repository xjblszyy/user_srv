package pg

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	. "user/utils"
	"user/utils/db/model"
)

var PgCnn *gorm.DB

func init(){
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
	PgCnn.AutoMigrate(&model.User{})
}




