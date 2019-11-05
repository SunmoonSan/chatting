/*
@desc : Created by San on 2019/11/1 17:14
*/
package ctrl

import (
	"chatting/model"
	"chatting/service"
	"chatting/util"
	"net/http"
)

var userService service.UserService

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	user, err := userService.Login(mobile, passwd)

	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOK(writer, user, "")
	}
}

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	util.Bind(request, &user)
	user, err := userService.Register(user.Mobile, user.Passwd, user.Nickname, user.Avatar, user.Sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOK(writer, user, "")
	}
}
