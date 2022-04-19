package main

import (
	"github.com/pp553933054/micro-go-book/ch13-seckill/pkg/bootstrap"
	conf "github.com/pp553933054/micro-go-book/ch13-seckill/pkg/config"
	"github.com/pp553933054/micro-go-book/ch13-seckill/pkg/mysql"
	"github.com/pp553933054/micro-go-book/ch13-seckill/sk-admin/setup"
)

func main() {
	mysql.InitMysql(conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.User, conf.MysqlConfig.Pwd, conf.MysqlConfig.Db) // conf.MysqlConfig.Db
	//setup.InitEtcd()
	setup.InitZk()
	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)

}
