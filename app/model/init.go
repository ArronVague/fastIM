package model

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	dsnName := "root:root@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, dsnName)
	if err != nil && err.Error() != "" {
		log.Fatal(err)
	}
	DbEngine.ShowSQL(true)
	//设置数据库连接数
	DbEngine.SetMaxOpenConns(10)
	//自动创建数据库
	if err = DbEngine.Sync(new(User), new(Community), new(Contact)); err != nil {
		panic(err)
	}

	fmt.Println("init database ok!")
}
