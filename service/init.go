/*
@desc : Created by San on 2019/10/31 00:58
*/
package service

import (
	"chatting/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngin *xorm.Engine
var err error

func init() {
	fmt.Println("init...")
	drivename := "mysql"
	DsName := "root:root@(127.0.0.1:3306)/chat?charset=utf8"
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if err != nil && "" != err.Error() {
		log.Fatal(err.Error())
	}
	DbEngin.ShowSQL(true)
	DbEngin.SetMaxOpenConns(2)
	fmt.Println("hello world")
	DbEngin.Sync(new(model.User))
	fmt.Println("init database ok...")
}
