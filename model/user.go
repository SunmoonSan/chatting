/*
@desc : Created by San on 2019/10/30 22:55
*/
package model

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MEN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	// 用户ID
	Id int64 `xorm: "pk autoincr bigint(20)" form:"id" json:"id"`
	// 手机号码
	Mobile string `xorm: "varchar(20)" form:"mobile" json:"mobile"`
	// 密码=plainpwd+salt, MD5
	Passwd string `xorm: "varchar(40)" form: "passwd" json:"-"`
	// 头像
	Avatar   string `xorm: "varchar(150)" from: "avatar" json:"avatar"`
	Sex      string `xorm: "varchar(2)" form: "sex" json:"sex"`
	Nickname string `xorm: "varchar(20)" form: "nickname" json:"nickname"`
	// 随机数
	Salt   string `xorm: "varchar(10)" form: "salt", json:"-"`
	Online int    `xorm: "int(10)" form: "online", json:"online"`
	Token  string `xorm: "varchar(40)" form: "token" json:"token"`
	Memo   string `xorm: varchar(140) form: "memo" json:"memo"`
	// 统计用户每天增量
	Createat time.Time `xorm: "datetime" from: "createat" json:"createat"`
}
