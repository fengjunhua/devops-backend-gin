package v1

import (
	"github/fengjunhua/devops-backend-gin/core/initialize"
	"github/fengjunhua/devops-backend-gin/models/mysql"
)

var mysqlClient = initialize.MysqlClient

func Register(u mysql.SysUser) (err error,userInter mysql.SysUser){
	var user mysql.SysUser

    mysqlClient.Create(&user)

	return nil,user
}

func GetUserById(){


}