/*
@desc : Created by San on 2019/10/31 00:53
*/
package service

import (
	"chatting/module"
	"chatting/util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
}

func (s *UserService) Register(mobile, plainpwd, nicknane, avatar, sex string) (user module.User, err error) {
	tmp := module.User{}
	_, err = DbEngin.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}

	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}

	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nicknane
	tmp.Sex = sex

	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Createat = time.Now()

	_, err = DbEngin.InsertOne(&tmp)
	return user, nil

}

func (s *UserService) Login(mobile, plainpwd string) (user module.User, err error) {

}
