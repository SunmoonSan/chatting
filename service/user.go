/*
@desc : Created by San on 2019/10/31 00:53
*/
package service

import (
	"chatting/model"
	"chatting/util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
}

func (s *UserService) Register(mobile, plainpwd, nicknane, avatar, sex string) (user model.User, err error) {
	tmp := model.User{}
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

	// token 可以是一个随机数
	tmp.Token = fmt.Sprintf("%08", rand.Int31())

	_, err = DbEngin.InsertOne(&tmp)
	return user, err
}

func (s *UserService) Login(
	mobile, //手机
	plainpwd string) (user model.User, err error) {

	//首先通过手机号查询用户
	tmp := model.User{}
	DbEngin.Where("mobile = ?", mobile).Get(&tmp)
	//如果没有找到
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	//查询到了比对密码
	if !util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}
	//刷新token,安全
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	//返回数据
	DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp, nil
}
