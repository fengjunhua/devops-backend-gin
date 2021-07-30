package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	MysqlClient *gorm.DB
)

func InitMysqlClient(){

	var err error
	dsn := "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	MysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

}

//无需close
func CloseMysqlClient(){


}